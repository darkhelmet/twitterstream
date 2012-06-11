package twitterstream_test

import (
    "encoding/json"
    ts "github.com/darkhelmet/twitterstream"
    . "launchpad.net/gocheck"
    "reflect"
    "testing"
    "time"
)

type StructTest struct {
    Struct       interface{}
    JsonField    string
    JsonValues   interface{}
    StructField  string
    StructValues interface{}
}

func Test(t *testing.T) { TestingT(t) }

type S struct{}

var _ = Suite(&S{})

func (s *S) TestEverything(c *C) {
    for _, test := range tableTests {
        jsonValues := reflect.ValueOf(test.JsonValues)
        structValues := reflect.ValueOf(test.StructValues)
        st := reflect.TypeOf(test.Struct)

        count := jsonValues.Len()
        if count != structValues.Len() {
            c.Errorf("must use same size slices for value pairs. %d != %d, %#v vs %#v", count, structValues.Len(), jsonValues.Interface(), structValues.Interface()) //, test.JsonValues, test.StructValues)
        }

        c.Logf("Encoding %s of %T", test.StructField, test.Struct)
        for index := 0; index < count; index += 1 {
            s := reflect.New(st).Elem()
            field := s.FieldByName(test.StructField)
            value := structValues.Index(index)
            if !field.CanSet() {
                c.Fatalf("Can't set %#v but should be able to", test.StructField)
            }
            field.Set(value)
            data := decodeJson(marshal(s.Interface()))
            expected := jsonValues.Index(index)
            if expected.Kind() == reflect.Ptr && expected.IsNil() {
                c.Assert(data[test.JsonField], IsNil)
            } else {
                c.Assert(data[test.JsonField], DeepEquals, convert(expected))
            }
        }

        c.Logf("Decoding %s of %T", test.StructField, test.Struct)
        for index := 0; index < count; index += 1 {
            j := make(map[string]interface{})
            value := reflect.ValueOf(&j).Elem()
            valueToSet := jsonValues.Index(index)
            key := reflect.ValueOf(test.JsonField)
            value.SetMapIndex(key, valueToSet)
            expected := structValues.Index(index)
            s := reflect.New(st)
            decode(s.Interface(), makeJson(j))
            s = s.Elem()
            field := s.FieldByName(test.StructField)
            if expected.Kind() == reflect.Ptr && expected.IsNil() {
                c.Assert(field.Interface(), IsNil)
            } else {
                c.Assert(convert(field), DeepEquals, convert(expected))
            }
        }
    }
}

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

func intaddr(value int) *int {
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
