package twitterstream

import (
    "encoding/json"
)

type IndexPair struct {
    Start int
    End   int
}

func (i IndexPair) MarshalJSON() ([]byte, error) {
    pair := []int{i.Start, i.End}
    return json.Marshal(pair)
}

func (i *IndexPair) UnmarshalJSON(data []byte) error {
    pair := [2]int{}
    err := json.Unmarshal(data, &pair)
    if err != nil {
        return err
    }
    i.Start = pair[0]
    i.End = pair[1]
    return nil
}
