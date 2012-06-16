package twitterstream

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "net/url"
    "strings"
    "time"
)

const (
    FilterUrl      = "https://stream.twitter.com/1/statuses/filter.json"
    DefaultTimeout = 1 * time.Minute
)

type Client struct {
    Username string
    Password string
    Timeout  time.Duration
}

func makeUrl(action, args string) string {
    return fmt.Sprintf("%s?%s=%s", FilterUrl, url.QueryEscape(action), url.QueryEscape(args))
}

func NewClient(username, password string) *Client {
    return NewClientTimeout(username, password, DefaultTimeout)
}

func NewClientTimeout(username, password string, timeout time.Duration) *Client {
    return &Client{
        Username: username,
        Password: password,
        Timeout:  timeout,
    }
}

func (c *Client) Track(keywords ...string) (*Connection, error) {
    uri := makeUrl("track", strings.Join(keywords, ","))
    req, err := http.NewRequest("POST", uri, nil)
    if err != nil {
        return nil, fmt.Errorf("twitterstream: Creating track request failed: %s", err)
    }

    req.SetBasicAuth(c.Username, c.Password)

    conn := newConnection(c.Timeout)
    resp, err := conn.client.Do(req)
    if err != nil {
        return nil, fmt.Errorf("twitterstream: Making track request failed: %s", err)
    }

    if resp.StatusCode != 200 {
        body, _ := ioutil.ReadAll(resp.Body)
        resp.Body.Close()
        return nil, fmt.Errorf("twitterstream: Track request failed (%d): %s", resp.StatusCode, body)
    }

    conn.setup(resp.Body)

    return conn, nil
}

func (c *Client) Follow(userIds ...string) (*Connection, error) {
    uri := makeUrl("follow", strings.Join(userIds, ","))
    req, err := http.NewRequest("POST", uri, nil)
    if err != nil {
        return nil, fmt.Errorf("twitterstream: Creating follow request failed: %s", err)
    }

    req.SetBasicAuth(c.Username, c.Password)

    conn := newConnection(c.Timeout)
    resp, err := conn.client.Do(req)
    if err != nil {
        return nil, fmt.Errorf("twitterstream: Making follow request failed: %s", err)
    }

    if resp.StatusCode != 200 {
        body, _ := ioutil.ReadAll(resp.Body)
        resp.Body.Close()
        return nil, fmt.Errorf("twitterstream: Track follow failed (%d): %s", resp.StatusCode, body)
    }

    conn.setup(resp.Body)

    return conn, nil
}
