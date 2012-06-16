package twitterstream

import (
    "encoding/json"
    "io"
    "net"
    "net/http"
    "time"
)

const DialTimeout = 5 * time.Second

type Connection struct {
    decoder *json.Decoder
    conn    net.Conn
    client  *http.Client
    closer  io.Closer
    timeout time.Duration
}

func (c *Connection) Close() error {
    // Have to close the raw connection, since closing the response body reader
    // will make Go try to read the request, which goes on forever.
    if err := c.conn.Close(); err != nil {
        c.closer.Close()
        return err
    }
    return c.closer.Close()
}

func (c *Connection) Next() (*Tweet, error) {
    var tweet Tweet
    c.conn.SetReadDeadline(time.Now().Add(c.timeout))
    if err := c.decoder.Decode(&tweet); err != nil {
        return nil, err
    }
    return &tweet, nil
}

func (c *Connection) setup(rc io.ReadCloser) {
    c.closer = rc
    c.decoder = json.NewDecoder(rc)
}

func newConnection(timeout time.Duration) *Connection {
    conn := &Connection{timeout: timeout}
    dialer := func(netw, addr string) (net.Conn, error) {
        netc, err := net.DialTimeout(netw, addr, DialTimeout)
        if err != nil {
            return nil, err
        }
        conn.conn = netc
        return netc, nil
    }

    conn.client = &http.Client{
        Transport: &http.Transport{
            Dial: dialer,
        },
    }

    return conn
}
