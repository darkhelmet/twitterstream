package twitterstream_test

import (
    "encoding/json"
    ts "github.com/darkhelmet/twitterstream"
    . "launchpad.net/gocheck"
    "testing"
    "time"
)

type JSON map[string]interface{}

func Test(t *testing.T) { TestingT(t) }

type S struct{}

var _ = Suite(&S{})

func marshal(v interface{}) []byte {
    data, err := json.Marshal(v)
    if err != nil {
        panic(err)
    }
    return data
}

func makeJson(v JSON) []byte {
    v["id"] = int64(12345)
    v["id_str"] = "12345"
    return marshal(v)
}

func decodeTweet(data []byte) ts.Tweet {
    var tweet ts.Tweet
    if err := json.Unmarshal(data, &tweet); err != nil {
        panic(err)
    }
    return tweet
}

func encodeTweet(tweet ts.Tweet) []byte {
    return marshal(tweet)
}

func decodeJson(data []byte) JSON {
    var v JSON
    if err := json.Unmarshal(data, &v); err != nil {
        panic(err)
    }
    return v
}

func (s *S) TestDecodeIds(c *C) {
    tweet := decodeTweet(makeJson(JSON{}))
    c.Assert(tweet.Id, Equals, int64(12345))
    c.Assert(tweet.IdString, Equals, "12345")
}

func (s *S) TestDecodeContributors(c *C) {
    tweet := decodeTweet(makeJson(JSON{
        "contributors": []JSON{
            JSON{
                "id":          int64(10),
                "id_str":      "10",
                "screen_name": "testing",
            },
        },
    }))
    c.Assert(tweet.Contributors[0], Equals, ts.Contributor{int64(10), "10", "testing"})
}

func (s *S) TestDecodeNullContributors(c *C) {
    tweet := decodeTweet(makeJson(JSON{"contributors": nil}))
    c.Assert(len(tweet.Contributors), Equals, 0)
}

func (s *S) TestEncodeNoContributors(c *C) {
    var tweet ts.Tweet
    c.Assert(len(tweet.Contributors), Equals, 0)
    data := decodeJson(encodeTweet(tweet))
    c.Assert(data["contributors"], IsNil)
}

func (s *S) TestDecodeCoordinates(c *C) {
    tweet := decodeTweet(makeJson(JSON{
        "coordinates": JSON{
            "coordinates": []float64{12.5, 10.2},
            "type":        "Point",
        },
    }))
    c.Assert(tweet.Coordinates, NotNil)
    c.Assert(tweet.Coordinates.Lat, Equals, ts.Latitude(10.2))
    c.Assert(tweet.Coordinates.Long, Equals, ts.Longitude(12.5))
    c.Assert(tweet.Coordinates.Type, Equals, "Point")
}

func (s *S) TestDecodeNoCoordinates(c *C) {
    tweet := decodeTweet(makeJson(JSON{"coordinates": nil}))
    c.Assert(tweet.Coordinates, IsNil)
}

func (s *S) TestEncodeCoordinates(c *C) {
    var tweet ts.Tweet
    tweet.Coordinates = ts.NewCoordinates(ts.Latitude(10.1), ts.Longitude(-12.5))
    data := decodeJson(encodeTweet(tweet))
    coords := data["coordinates"].(map[string]interface{})
    points := coords["coordinates"].([]interface{})
    c.Assert(points[0], Equals, -12.5)
    c.Assert(points[1], Equals, 10.1)
}

func (s *S) TestEncodeNoCoordinates(c *C) {
    var tweet ts.Tweet
    c.Assert(tweet.Coordinates, IsNil)
    data := decodeJson(encodeTweet(tweet))
    c.Assert(data["coordinates"], IsNil)
}

func (s *S) TestDecodeCreatedAT(c *C) {
    tweet := decodeTweet(makeJson(JSON{"created_at": "Wed Aug 27 13:08:45 +0000 2008"}))
    c.Assert(tweet.CreatedAt.Time.Day(), Equals, 27)
}

func (s *S) TestEncodeCreatedAt(c *C) {
    timeString := "Wed Aug 27 13:08:45 +0000 2008"
    var tweet ts.Tweet
    tweet.CreatedAt.Time, _ = time.Parse(ts.TimeFormat, timeString)
    data := decodeJson(encodeTweet(tweet))
    c.Assert(data["created_at"], Equals, timeString)
}
