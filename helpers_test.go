package twitterstream_test

import (
    ts "github.com/darkhelmet/twitterstream"
    "time"
)

func float64s(is ...interface{}) []interface{} {
    values := make([]interface{}, 0, len(is))
    for _, i := range is {
        switch v := i.(type) {
        case nil:
            values = append(values, nil)
        case int:
            values = append(values, float64(v))
        default:
            panic("add things")
        }
    }
    return values
}

func booladdr(value bool) *bool {
    return &value
}

func int64addr(value int64) *int64 {
    return &value
}

func straddr(value string) *string {
    return &value
}

func mustParseTime(value string) ts.Time {
    var t ts.Time
    var err error
    t.Time, err = time.Parse(ts.TimeFormat, value)
    if err != nil {
        panic(err)
    }
    return t
}

func marshal(v interface{}) []byte {
    data, err := json.Marshal(v)
    if err != nil {
        panic(err)
    }
    return data
}

func makeJson(hash JSON) []byte {
    return marshal(hash)
}

func decodeTweet(data []byte) ts.Tweet {
    var t ts.Tweet
    if err := json.Unmarshal(data, &t); err != nil {
        panic(err)
    }
    return t
}

func decode(i interface{}, data []byte) {
    if err := json.Unmarshal(data, i); err != nil {
        panic(err)
    }
}

func decodeJson(data []byte) JSON {
    var v JSON
    if err := json.Unmarshal(data, &v); err != nil {
        panic(err)
    }
    return v
}

func convert(value reflect.Value) interface{} {
    if value.Kind() == reflect.Ptr {
        return value.Elem().Interface()
    }
    return value.Interface()
}
