package twitterstream

import (
    "bufio"
    // "os"
    "encoding/json"
    "io"
    "net"
    "net/http"
    "time"
)

const DialTimeout = 5 * time.Second

type Connection struct {
    decoder    *json.Decoder
    httpConn   net.Conn
    httpClient *http.Client
    closer     io.Closer
}

func (c *Connection) Close() error {
    if err := c.httpConn.Close(); err != nil {
        return err
    }
    return c.closer.Close()
}

func (c *Connection) Next() (*Tweet, error) {
    var tweet Tweet
    if err := c.decoder.Decode(&tweet); err != nil {
        return nil, err
    }
    return &tweet, nil
}

func (c *Connection) setup(rc io.ReadCloser) {
    // rd, wr := io.Pipe()
    // mwr := io.MultiWriter(os.Stdout, wr)
    c.closer = rc
    c.decoder = json.NewDecoder(bufio.NewReader(rc))
    // go io.Copy(mwr, rc)
}

func newConnection() *Connection {
    conn := new(Connection)
    dialer := func(netw, addr string) (net.Conn, error) {
        netc, err := net.DialTimeout(netw, addr, DialTimeout)
        if err != nil {
            return nil, err
        }
        conn.httpConn = netc
        return netc, nil
    }

    httpClient := &http.Client{
        Transport: &http.Transport{
            Dial: dialer,
        },
    }
    conn.httpClient = httpClient
    return conn
}
