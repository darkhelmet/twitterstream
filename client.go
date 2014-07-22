package twitterstream

import (
	"fmt"
	"github.com/garyburd/go-oauth/oauth"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	FilterUrl      = "https://stream.twitter.com/1.1/statuses/filter.json"
	DefaultTimeout = 1 * time.Minute
)

type Client struct {
	Oauth       *oauth.Client
	Credentials *oauth.Credentials
	Timeout     time.Duration
}

func NewClient(consumerKey, consumerSecret, accessToken, accessSecret string) *Client {
	return NewClientTimeout(consumerKey, consumerSecret, accessToken, accessSecret, DefaultTimeout)
}

func NewClientTimeout(consumerKey, consumerSecret, accessToken, accessSecret string, timeout time.Duration) *Client {
	return &Client{
		Oauth: &oauth.Client{
			Credentials: oauth.Credentials{
				Token:  consumerKey,
				Secret: consumerSecret,
			},
		},
		Credentials: &oauth.Credentials{
			Token:  accessToken,
			Secret: accessSecret,
		},
		Timeout: timeout,
	}
}

func (c *Client) Track(keywords ...string) (*Connection, error) {
	form := url.Values{"track": {strings.Join(keywords, ",")}}
	return c.filter(form)
}

func (c *Client) Follow(userIds ...string) (*Connection, error) {
	form := url.Values{"follow": {strings.Join(userIds, ",")}}
	return c.filter(form)
}

func (c *Client) Locations(p1 Point, p2 Point) (*Connection, error) {
	form := url.Values{"locations": {fmt.Sprintf("%f,%f,%f,%f", p1.Long, p1.Lat, p2.Long, p2.Lat)}}
	return c.filter(form)
}

func (c *Client) Sample() (*Connection, error) {
	req, err := http.NewRequest("GET", "https://stream.twitter.com/1.1/statuses/sample.json", nil)
	if err != nil {
		return nil, fmt.Errorf("twitterstream: creating sample request failed: %s", err)
	}

	req.Header.Set("Authorization", c.Oauth.AuthorizationHeader(c.Credentials, "GET", req.URL, nil))

	conn := newConnection(c.Timeout)
	resp, err := conn.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("twitterstream: making sample request failed: %s", err)
	}

	if resp.StatusCode != 200 {
		body, _ := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		return nil, fmt.Errorf("twitterstream: sample failed (%d): %s", resp.StatusCode, body)
	}

	conn.setup(resp.Body)

	return conn, nil
}

func (c *Client) filter(form url.Values) (*Connection, error) {
	req, err := http.NewRequest("POST", FilterUrl, strings.NewReader(form.Encode()))
	if err != nil {
		return nil, fmt.Errorf("twitterstream: creating filter request failed: %s", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", c.Oauth.AuthorizationHeader(c.Credentials, "POST", req.URL, form))

	conn := newConnection(c.Timeout)
	resp, err := conn.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("twitterstream: making filter request failed: %s", err)
	}

	if resp.StatusCode != 200 {
		body, _ := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		return nil, fmt.Errorf("twitterstream: filter failed (%d): %s", resp.StatusCode, body)
	}

	conn.setup(resp.Body)

	return conn, nil
}
