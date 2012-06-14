# twitterstream

Twitter streaming API for Go. [Documentation](http://go.pkgdoc.org/github.com/darkhelmet/twitterstream)

    package main

    import (
        "log"
        "time"
        "github.com/darkhelmet/twitterstream"
    )

    func decode(conn *twitterstream.Connection) {
        for {
            if tweet, err := conn.Next(); err == nil {
                log.Println("%s said: %s", tweet.User.ScreenName, tweet.Text)
            } else {
                log.Printf("Failed decoding tweet: %s", err)
                return
            }
        }
    }

    func main() {
        client := twitterstream.NewClient("username", "password")
        for {
            conn, err := client.Track("Justin Bieber,American Idol")
            if err != nil {
                log.Println("Tracking failed, sleeping for 1 minute")
                time.Sleep(1 * time.Minute)
                continue
            }
            decode(conn)
        }
    }

# License

Copyright © 2012 Daniel Huckstep. See LICENSE for details.
