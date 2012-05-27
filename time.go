package twitterstream

import (
    "encoding/json"
    "time"
)

const (
    TimeFormat = "Mon Jan _2 15:04:05 +0000 2006"
)

type Time struct {
    Time time.Time
}

func (t Time) MarshalJSON() ([]byte, error) {
    formatted := t.Time.Format(TimeFormat)
    return json.Marshal(formatted)
}

func (t *Time) UnmarshalJSON(data []byte) error {
    var s string
    err := json.Unmarshal(data, &s)
    if err != nil {
        return err
    }
    t.Time, err = time.Parse(TimeFormat, s)
    return err
}
