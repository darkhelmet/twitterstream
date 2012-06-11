package twitterstream_test

import (
    ts "github.com/darkhelmet/twitterstream"
    . "launchpad.net/gocheck"
)

type JSON map[string]interface{}

func (s *S) TestDecodeNullContributors(c *C) {
    tweet := decodeTweet(makeJson(JSON{"contributors": nil}))
    c.Assert(len(tweet.Contributors), Equals, 0)
}

func (s *S) TestEncodeNoContributors(c *C) {
    var tweet ts.Tweet
    c.Assert(len(tweet.Contributors), Equals, 0)
    data := decodeJson(marshal(tweet))
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
}

func (s *S) TestEncodeCoordinates(c *C) {
    var tweet ts.Tweet
    tweet.Coordinates = &ts.Point{ts.Latitude(10.1), ts.Longitude(-12.5)}
    data := decodeJson(marshal(tweet))
    coords := data["coordinates"].(map[string]interface{})
    points := coords["coordinates"].([]interface{})
    c.Assert(points[0], Equals, -12.5)
    c.Assert(points[1], Equals, 10.1)
}

func (s *S) TestDecodePlace(c *C) {
    tweet := decodeTweet(makeJson(JSON{
        "place": JSON{
            "attributes": JSON{
                "street_address": "123 Sesame St",
            },
            "bounding_box": JSON{
                "coordinates": []interface{}{
                    []interface{}{
                        []float64{1.2, -3.4},
                        []float64{5.6, 7.8},
                        []float64{-9.0, -1.1},
                        []float64{2.2, -3.3},
                    },
                },
                "type": "Polygon",
            },
        },
    }))
    p := tweet.Place
    c.Assert(p.Attributes["street_address"], Equals, "123 Sesame St")
    c.Assert(p.BoundingBox.Points[0].Lat.Float64(), Equals, -3.4)
    c.Assert(p.BoundingBox.Points[0].Long.Float64(), Equals, 1.2)
    c.Assert(p.BoundingBox.Points[2].Long.Float64(), Equals, -9.0)
    c.Assert(p.BoundingBox.Points[3].Lat.Float64(), Equals, -3.3)
}

func (s *S) TestEncodePlace(c *C) {
    var tweet ts.Tweet
    tweet.Place = &ts.Place{
        Attributes: map[string]interface{}{
            "foo": "bar",
        },
        BoundingBox: ts.Box{
            Points: []ts.Point{
                ts.Point{1.0, 2.5},
                ts.Point{3.4, 4.1},
            },
        },
    }
    data := decodeJson(marshal(tweet))
    place := data["place"].(map[string]interface{})

    box := place["bounding_box"].(map[string]interface{})
    c.Assert(box["type"], Equals, "Polygon")

    points := box["coordinates"].([]interface{})[0].([]interface{})
    p1 := points[0].([]interface{})
    p2 := points[1].([]interface{})
    c.Assert(p1[0], Equals, 2.5)
    c.Assert(p1[1], Equals, 1.0)
    c.Assert(p2[0], Equals, 4.1)
    c.Assert(p2[1], Equals, 3.4)
}
