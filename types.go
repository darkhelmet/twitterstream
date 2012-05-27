package twitterstream

import ()

type Indices struct {
    Start int
    End   int
}

type Hashtag struct {
    Text    string  `json:"text"`
    Indices Indices `json:"indices"`
}

type Medium struct {
}

type Entities struct {
    Hashtags []Hashtag
    Media    []Medium
}

type Place struct {
}

type Contributor struct {
    Id         int64  `json:"id"`
    IdString   string `json:"id_str"`
    ScreenName string `json:"screen_name"`
}

type Tweet struct {
    // The integer representation of the unique identifier for this Tweet. This number is greater than 53 bits and some programming languages may have difficulty/silent defects in interpreting it. Using a signed 64 bit integer for storing this identifier is safe. Use id_str for fetching the identifier to stay on the safe side. See Twitter IDs, JSON and Snowflake.
    Id  int64 `json:"id"`

    // The string representation of the unique identifier for this Tweet. Implementations should use this rather than the large integer in id.
    IdString string `json:"id_str"`

    // An collection of brief user objects (usually only one) indicating users who contributed to the authorship of the tweet, on behalf of the official tweet author.
    Contributors []Contributor `json:"contributors"`

    // Represents the geographic location of this Tweet as reported by the user or client application.
    Coordinates *Coordinates `json:"coordinates"`

    // Time when this Tweet was created.
    CreatedAt Time `json:"created_at"`
}
