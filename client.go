package twitterstream

import (
    "compress/gzip"
    "fmt"
    "io"
    "io/ioutil"
    "net/http"
    "net/url"
    "strings"
)

const (
    FilterUrl = "https://stream.twitter.com/1/statuses/filter.json"
)

type Client struct {
    Username string
    Password string
}

func NewClient(username, password string) *Client {
    return &Client{
        Username: username,
        Password: password,
    }
}

func (c *Client) Track(keywords ...string) (*Connection, error) {
    uri := fmt.Sprintf("%s?track=%s", FilterUrl, url.QueryEscape(strings.Join(keywords, ",")))
    req, err := http.NewRequest("POST", uri, nil)
    if err != nil {
        return nil, fmt.Errorf("twitterstream: Creating the request failed: %s", err)
    }

    req.SetBasicAuth(c.Username, c.Password)
    req.Header.Add("Accept-Encoding", "gzip")

    conn := newConnection()
    resp, err := conn.httpClient.Do(req)
    if err != nil {
        return nil, fmt.Errorf("twitterstream: Making the request failed: %s", err)
    }

    if resp.StatusCode != 200 {
        body, _ := ioutil.ReadAll(resp.Body)
        resp.Body.Close()
        return nil, fmt.Errorf("twitterstream: Request failed (%d): %s", resp.StatusCode, body)
    }

    var rc io.ReadCloser = resp.Body
    if resp.Header.Get("Content-Encoding") == "gzip" {
        rc, err = gzip.NewReader(rc)
        if err != nil {
            return nil, err
        }
    }

    conn.setup(rc)

    return conn, nil
}
