package main

import (
    "code.google.com/p/go.net/websocket"
    "flag"
    "github.com/darkhelmet/twitterstream"
    T "html/template"
    "io"
    "log"
    "net/http"
    "net/url"
    "os"
    "strings"
    "time"
)

type subscription struct {
    conn *websocket.Conn
    ch   chan bool
    s    bool
}

var (
    home = T.Must(T.New("home").Parse(`
        <html>
        <head>
            <title>Twitter Image Stream</title>
            <style type="text/css">
                body {
                    width: 600px;
                    margin: 0px auto;
                    text-align: center;
                }

                img {
                    max-width: 100%;
                    border-radius: 15px;
                    margin: 5px;
                }
            </style>
            <script type="text/javascript" src="https://ajax.googleapis.com/ajax/libs/jquery/2.0.3/jquery.min.js"></script>
            <script type="text/javascript">
                var images = [];
                setInterval(function() {
                    var img = images.shift();
                    if (img) {
                        $(img).hide().
                            css('max-height', $(window).height() + 'px').
                            prependTo($('body')).
                            fadeIn();
                    }
                }, 5000);
                var conn = new WebSocket("ws://{{.}}/ws");
                conn.onclose = function(event) {
                    console.log('closed');
                };
                conn.onmessage = function(event) {
                    var img = new Image();
                    img.onload = function() {
                        images.push(img);
                    };
                    img.src = event.data;
                };
            </script>
        </head>
        <body>
        </body>
        </html>
    `))

    consumerKey    = flag.String("consumer-key", "", "consumer key")
    consumerSecret = flag.String("consumer-secret", "", "consumer secret")
    accessToken    = flag.String("access-token", "", "access token")
    accessSecret   = flag.String("access-secret", "", "access token secret")
    keywords       = flag.String("keywords", "", "keywords to track")
    listen         = flag.String("listen", ":8080", "Spec to listen on")
    subscriptions  = make(chan subscription)
    messages       = make(chan []byte, 5)
    wait           = 1
    maxWait        = 600 // Seconds
)

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func init() {
    flag.Parse()

    if *consumerKey == "" || *consumerSecret == "" {
        log.Fatalln("consumer tokens left blank")
    }

    if *accessToken == "" || *accessSecret == "" {
        log.Fatalln("access tokens left blank")
    }

    if *keywords == "" {
        log.Fatalln("keywords left blank")
    }

    go hub()
    go downloader()
}

func last(s []string) string {
    if len(s) == 0 {
        return ""
    }
    return s[len(s)-1]
}

func download(uri string) bool {
    u, err := url.Parse(uri)
    if err != nil {
        log.Printf("failed parsing URI: %s", err)
        return false
    }
    filename := last(strings.Split(u.Path, "/"))

    file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0644)
    if err != nil {
        log.Printf("error opening file: %s", err)
        return false
    }
    defer file.Close()

    log.Printf("downloading %s", uri)

    if resp, err := http.Get(uri); err != nil {
        log.Printf("failed downloading: %s", err)
        return false
    } else {
        defer resp.Body.Close()
        io.Copy(file, resp.Body)
    }
    return true
}

func decodeTweets(conn *twitterstream.Connection) {
    for {
        if tweet, err := conn.Next(); err == nil {
            for _, medium := range tweet.Entities.Media {
                if download(medium.MediaUrl) {
                    messages <- []byte(medium.MediaUrl)
                }
            }
        } else {
            log.Printf("decoding tweet failed: %s", err)
            conn.Close()
            return
        }
    }
}

func downloader() {
    client := twitterstream.NewClient(*consumerKey, *consumerSecret, *accessToken, *accessSecret)
    for {
        log.Printf("tracking keywords %s", *keywords)
        conn, err := client.Track(*keywords)
        if err != nil {
            log.Printf("tracking failed: %s", err)
            wait = wait << 1
            log.Printf("waiting for %d seconds before reconnect", min(wait, maxWait))
            time.Sleep(time.Duration(min(wait, maxWait)) * time.Second)
            continue
        } else {
            wait = 1
        }
        decodeTweets(conn)
    }
}

func hub() {
    conns := make(map[*websocket.Conn]chan bool)
    for {
        select {
        case sub := <-subscriptions:
            if sub.s {
                conns[sub.conn] = sub.ch
            } else {
                delete(conns, sub.conn)
            }
        case message := <-messages:
            for conn, ch := range conns {
                if _, err := conn.Write(message); err != nil {
                    conn.Close()
                    ch <- false
                    close(ch)
                }
            }
        }
    }
}

func homeHandler(w http.ResponseWriter, req *http.Request) {
    home.Execute(w, req.Host)
}

func wsHandler(ws *websocket.Conn) {
    ch := make(chan bool)
    subscriptions <- subscription{ws, ch, true}
    subscriptions <- subscription{ws, ch, <-ch}
}

func main() {
    http.HandleFunc("/", homeHandler)
    http.Handle("/ws", websocket.Handler(wsHandler))
    err := http.ListenAndServe(*listen, nil)
    if err != nil {
        log.Fatalf("failed to listen: %s", err)
    }
}
