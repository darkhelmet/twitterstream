package twitterstream

import (
    "encoding/json"
    "errors"
    "fmt"
)

var (
    MissingCoordinates = errors.New("Missing coordinates")
)

type Latitude float64
type Longitude float64

func (l Latitude) Float64() float64 {
    return float64(l)
}

func (l Longitude) Float64() float64 {
    return float64(l)
}

type Point struct {
    Lat  Latitude
    Long Longitude
}

type rawPoint struct {
    Points [2]float64 `json:"coordinates"`
    Type   string     `json:"type"`
}

func (c *Point) UnmarshalJSON(data []byte) error {
    var rc rawPoint
    err := json.Unmarshal(data, &rc)
    if err != nil {
        return err
    }
    c.Long = Longitude(rc.Points[0])
    c.Lat = Latitude(rc.Points[1])
    return nil
}

func (c *Point) MarshalJSON() ([]byte, error) {
    var rc rawPoint
    rc.Type = "Point"
    rc.Points[0] = float64(c.Long)
    rc.Points[1] = float64(c.Lat)
    return json.Marshal(rc)
}

type Box struct {
    Points []Point
}

func (b *Box) UnmarshalJSON(data []byte) (err error) {
    var v map[string]interface{}
    err = json.Unmarshal(data, &v)
    if err != nil {
        return
    }

    defer func() {
        if r := recover(); r != nil {
            err = fmt.Errorf("%s: failed decoding %T from %#v", r, b, string(data))
        }
    }()

    inner := v["coordinates"].([]interface{})[0].([]interface{})
    for _, array := range inner {
        pair := array.([]interface{})
        point := Point{Latitude(pair[1].(float64)), Longitude(pair[0].(float64))}
        b.Points = append(b.Points, point)
    }

    return nil
}

func (b *Box) MarshalJSON() ([]byte, error) {
    coords := make([]interface{}, 0, len(b.Points))
    for _, point := range b.Points {
        coords = append(coords, []interface{}{point.Long, point.Lat})
    }
    v := map[string]interface{}{
        "coordinates": []interface{}{coords},
        "type":        "Polygon",
    }
    return json.Marshal(v)
}
