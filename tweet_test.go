package twitterstream_test

import (
    "encoding/json"
    ts "github.com/darkhelmet/twitterstream"
    . "launchpad.net/gocheck"
    "reflect"
    "testing"
)

type JSON map[string]interface{}

type StructTest struct {
    Struct       interface{}
    JsonField    string
    JsonValues   interface{}
    StructField  string
    StructValues interface{}
}

var (
    tests = []StructTest{
        StructTest{
            Struct:       ts.Tweet{},
            JsonField:    "favorited",
            JsonValues:   []interface{}{nil, true, false},
            StructField:  "Favorited",
            StructValues: []*bool{nil, booladdr(true), booladdr(false)},
        },
        StructTest{
            Struct:       ts.Tweet{},
            JsonField:    "id",
            JsonValues:   float64s(1, 10, 25000),
            StructField:  "Id",
            StructValues: []int64{1, 10, 25000},
        },
        StructTest{
            Struct:       ts.Tweet{},
            JsonField:    "created_at",
            JsonValues:   []interface{}{"Wed Aug 27 13:08:45 +0000 2008"},
            StructField:  "CreatedAt",
            StructValues: []ts.Time{mustParseTime("Wed Aug 27 13:08:45 +0000 2008")},
        },
        StructTest{
            Struct:       ts.Tweet{},
            JsonField:    "in_reply_to_screen_name",
            JsonValues:   []interface{}{nil, "twitter", nil, "batman"},
            StructField:  "InReplyToScreenName",
            StructValues: []*string{nil, straddr("twitter"), nil, straddr("batman")},
        },
        StructTest{
            Struct:       ts.Tweet{},
            JsonField:    "in_reply_to_status_id",
            JsonValues:   float64s(nil, 123, 12345),
            StructField:  "InReplyToStatusId",
            StructValues: []*int64{nil, int64addr(123), int64addr(12345)},
        },
        StructTest{
            Struct:       ts.Tweet{},
            JsonField:    "in_reply_to_status_id_str",
            JsonValues:   []interface{}{nil, "123", "12345"},
            StructField:  "InReplyToStatusIdStr",
            StructValues: []*string{nil, straddr("123"), straddr("12345")},
        },
        StructTest{
            Struct:       ts.Tweet{},
            JsonField:    "in_reply_to_user_id",
            JsonValues:   float64s(nil, 123, 12345),
            StructField:  "InReplyToUserId",
            StructValues: []*int64{nil, int64addr(123), int64addr(12345)},
        },
        StructTest{
            Struct:       ts.Tweet{},
            JsonField:    "in_reply_to_user_id_str",
            JsonValues:   []interface{}{nil, "123", "12345"},
            StructField:  "InReplyToUserIdStr",
            StructValues: []*string{nil, straddr("123"), straddr("12345")},
        },
        StructTest{
            Struct:       ts.Tweet{},
            JsonField:    "possibly_sensitive",
            JsonValues:   []interface{}{nil, true, false},
            StructField:  "PossiblySensitive",
            StructValues: []*bool{nil, booladdr(true), booladdr(false)},
        },
        StructTest{
            Struct:       ts.Tweet{},
            JsonField:    "withheld_scope",
            JsonValues:   []interface{}{"", "status", "location"},
            StructField:  "WithheldScope",
            StructValues: []string{"", "status", "location"},
        },
        StructTest{
            Struct:       ts.Tweet{},
            JsonField:    "withheld_in_countries",
            JsonValues:   []interface{}{"", "US", "RU, CN, AU"},
            StructField:  "WithheldInCountries",
            StructValues: []string{"", "US", "RU, CN, AU"},
        },
        StructTest{
            Struct:       ts.Tweet{},
            JsonField:    "retweeted",
            JsonValues:   []interface{}{true, false},
            StructField:  "Retweeted",
            StructValues: []bool{true, false},
        },
        StructTest{
            Struct:       ts.Tweet{},
            JsonField:    "truncated",
            JsonValues:   []interface{}{true, false},
            StructField:  "Truncated",
            StructValues: []bool{true, false},
        },
        StructTest{
            Struct:       ts.Tweet{},
            JsonField:    "source",
            JsonValues:   []interface{}{"", "web"},
            StructField:  "Source",
            StructValues: []string{"", "web"},
        },
        StructTest{
            Struct:       ts.Tweet{},
            JsonField:    "text",
            JsonValues:   []interface{}{"", "wat", "superbatman and a bunch of stuff"},
            StructField:  "Text",
            StructValues: []string{"", "wat", "superbatman and a bunch of stuff"},
        },
        StructTest{
            Struct:       ts.Tweet{},
            JsonField:    "retweet_count",
            JsonValues:   float64s(0, 10, 2500),
            StructField:  "RetweetCount",
            StructValues: []int{0, 10, 2500},
        },
        StructTest{
            Struct:       ts.Tweet{},
            JsonField:    "place",
            JsonValues:   []interface{}{nil},
            StructField:  "Place",
            StructValues: []*ts.Place{nil},
        },
        StructTest{
            Struct:       ts.Tweet{},
            JsonField:    "coordinates",
            JsonValues:   []interface{}{nil},
            StructField:  "Coordinates",
            StructValues: []*ts.Point{nil},
        },
        StructTest{
            Struct:       ts.Contributor{},
            JsonField:    "id",
            JsonValues:   float64s(1, 10),
            StructField:  "Id",
            StructValues: []int64{1, 10},
        },
        StructTest{
            Struct:       ts.Place{},
            JsonField:    "country",
            JsonValues:   []interface{}{"Canada", "United States"},
            StructField:  "Country",
            StructValues: []string{"Canada", "United States"},
        },
        StructTest{
            Struct:       ts.Place{},
            JsonField:    "country_code",
            JsonValues:   []interface{}{"CA", "US"},
            StructField:  "CountryCode",
            StructValues: []string{"CA", "US"},
        },
        StructTest{
            Struct:       ts.Place{},
            JsonField:    "full_name",
            JsonValues:   []interface{}{"Canada", "United States"},
            StructField:  "FullName",
            StructValues: []string{"Canada", "United States"},
        },
        StructTest{
            Struct:       ts.Place{},
            JsonField:    "id",
            JsonValues:   []interface{}{"edmonton"},
            StructField:  "Id",
            StructValues: []string{"edmonton"},
        },
        StructTest{
            Struct:       ts.Place{},
            JsonField:    "name",
            JsonValues:   []interface{}{"San Diego", "Edmonton"},
            StructField:  "Name",
            StructValues: []string{"San Diego", "Edmonton"},
        },
        StructTest{
            Struct:       ts.Place{},
            JsonField:    "place_type",
            JsonValues:   []interface{}{"City", "Country"},
            StructField:  "Type",
            StructValues: []string{"City", "Country"},
        },
        StructTest{
            Struct:       ts.Place{},
            JsonField:    "url",
            JsonValues:   []interface{}{"http://api.twitter.com/1/geo/id/sandiego.json"},
            StructField:  "Url",
            StructValues: []string{"http://api.twitter.com/1/geo/id/sandiego.json"},
        },
        StructTest{
            Struct:       ts.Medium{},
            JsonField:    "id",
            JsonValues:   float64s(1, 22, 555),
            StructField:  "Id",
            StructValues: []int64{1, 22, 555},
        },
        StructTest{
            Struct:       ts.Medium{},
            JsonField:    "id_str",
            JsonValues:   []interface{}{"1", "22", "12345"},
            StructField:  "IdStr",
            StructValues: []string{"1", "22", "12345"},
        },
        StructTest{
            Struct:       ts.Medium{},
            JsonField:    "media_url",
            JsonValues:   []interface{}{"foo", "https://dev.twitter.com/docs/tweet-entities"},
            StructField:  "MediaUrl",
            StructValues: []string{"foo", "https://dev.twitter.com/docs/tweet-entities"},
        },
        StructTest{
            Struct:       ts.Medium{},
            JsonField:    "media_url_https",
            JsonValues:   []interface{}{"foo", "https://dev.twitter.com/docs/tweet-entities"},
            StructField:  "SecureMediaUrl",
            StructValues: []string{"foo", "https://dev.twitter.com/docs/tweet-entities"},
        },
        StructTest{
            Struct:       ts.Medium{},
            JsonField:    "url",
            JsonValues:   []interface{}{"foo", "https://dev.twitter.com/docs/tweet-entities"},
            StructField:  "Url",
            StructValues: []string{"foo", "https://dev.twitter.com/docs/tweet-entities"},
        },
        StructTest{
            Struct:       ts.Medium{},
            JsonField:    "display_url",
            JsonValues:   []interface{}{"foo", "https://dev.twitter.com/docs/tweet-entities"},
            StructField:  "DisplayUrl",
            StructValues: []string{"foo", "https://dev.twitter.com/docs/tweet-entities"},
        },
        StructTest{
            Struct:       ts.Medium{},
            JsonField:    "expanded_url",
            JsonValues:   []interface{}{"foo", "https://dev.twitter.com/docs/tweet-entities"},
            StructField:  "ExpandedUrl",
            StructValues: []string{"foo", "https://dev.twitter.com/docs/tweet-entities"},
        },
        StructTest{
            Struct:       ts.Medium{},
            JsonField:    "indices",
            JsonValues:   [][]interface{}{float64s(5, 10), float64s(0, 9)},
            StructField:  "Indices",
            StructValues: []ts.IndexPair{ts.IndexPair{5, 10}, ts.IndexPair{0, 9}},
        },
        StructTest{
            Struct:       ts.Size{},
            JsonField:    "w",
            JsonValues:   float64s(1, 10, 1000),
            StructField:  "Width",
            StructValues: []int{1, 10, 1000},
        },
        StructTest{
            Struct:       ts.Size{},
            JsonField:    "h",
            JsonValues:   float64s(1, 10, 1000),
            StructField:  "Height",
            StructValues: []int{1, 10, 1000},
        },
        StructTest{
            Struct:       ts.Size{},
            JsonField:    "resize",
            JsonValues:   []interface{}{"fit", "crop"},
            StructField:  "Resize",
            StructValues: []string{"fit", "crop"},
        },
        StructTest{
            Struct:       ts.Sizes{},
            JsonField:    "large",
            JsonValues:   []interface{}{map[string]interface{}{"w": float64(100), "h": float64(300), "resize": "fit"}},
            StructField:  "L",
            StructValues: []ts.Size{ts.Size{100, 300, "fit"}},
        },
        StructTest{
            Struct:       ts.Url{},
            JsonField:    "url",
            JsonValues:   []interface{}{"foo", "https://dev.twitter.com/docs/tweet-entities"},
            StructField:  "Url",
            StructValues: []string{"foo", "https://dev.twitter.com/docs/tweet-entities"},
        },
        StructTest{
            Struct:       ts.Url{},
            JsonField:    "display_url",
            JsonValues:   []interface{}{"foo", "https://dev.twitter.com/docs/tweet-entities"},
            StructField:  "DisplayUrl",
            StructValues: []string{"foo", "https://dev.twitter.com/docs/tweet-entities"},
        },
        StructTest{
            Struct:       ts.Url{},
            JsonField:    "expanded_url",
            JsonValues:   []interface{}{"foo", "https://dev.twitter.com/docs/tweet-entities"},
            StructField:  "ExpandedUrl",
            StructValues: []string{"foo", "https://dev.twitter.com/docs/tweet-entities"},
        },
        StructTest{
            Struct:       ts.Url{},
            JsonField:    "indices",
            JsonValues:   [][]interface{}{float64s(5, 10), float64s(0, 9)},
            StructField:  "Indices",
            StructValues: []ts.IndexPair{ts.IndexPair{5, 10}, ts.IndexPair{0, 9}},
        },
        StructTest{
            Struct:       ts.Mention{},
            JsonField:    "id",
            JsonValues:   float64s(1, 22, 555),
            StructField:  "Id",
            StructValues: []int64{1, 22, 555},
        },
        StructTest{
            Struct:       ts.Mention{},
            JsonField:    "id_str",
            JsonValues:   []interface{}{"1", "22", "12345"},
            StructField:  "IdStr",
            StructValues: []string{"1", "22", "12345"},
        },
        StructTest{
            Struct:       ts.Mention{},
            JsonField:    "screen_name",
            JsonValues:   []interface{}{"bill", "bob"},
            StructField:  "ScreenName",
            StructValues: []string{"bill", "bob"},
        },
        StructTest{
            Struct:       ts.Mention{},
            JsonField:    "full_name",
            JsonValues:   []interface{}{"Bruce Wayne"},
            StructField:  "FullName",
            StructValues: []string{"Bruce Wayne"},
        },
        StructTest{
            Struct:       ts.Mention{},
            JsonField:    "indices",
            JsonValues:   [][]interface{}{float64s(5, 10), float64s(0, 9)},
            StructField:  "Indices",
            StructValues: []ts.IndexPair{ts.IndexPair{5, 10}, ts.IndexPair{0, 9}},
        },
    }
)

func Test(t *testing.T) { TestingT(t) }

type S struct{}

var _ = Suite(&S{})

func (s *S) TestEverything(c *C) {
    for _, test := range tests {
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
