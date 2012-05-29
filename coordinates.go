package twitterstream

import (
    "encoding/json"
)

type Latitude float64
type Longitude float64

type Coordinates struct {
    Lat  Latitude
    Long Longitude
    Type string
}

func NewCoordinates(lat Latitude, long Longitude) *Coordinates {
    return &Coordinates{lat, long, "Point"}
}

type rawCoordinates struct {
    Points [2]float64 `json:"coordinates"`
    Type   string     `json:"type"`
}

func (c *Coordinates) UnmarshalJSON(data []byte) error {
    var rc rawCoordinates
    err := json.Unmarshal(data, &rc)
    if err != nil {
        return err
    }
    c.Lat = Latitude(rc.Points[1])
    c.Long = Longitude(rc.Points[0])
    c.Type = "Point"
    return nil
}

func (c *Coordinates) MarshalJSON() ([]byte, error) {
    var rc rawCoordinates
    rc.Type = c.Type
    rc.Points[0] = float64(c.Long)
    rc.Points[1] = float64(c.Lat)
    return json.Marshal(rc)
}
