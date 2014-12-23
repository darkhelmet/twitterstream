package twitterstream

type Hashtag struct {
    Text    string    `json:"text"`
    Indices IndexPair `json:"indices"`
}

type Size struct {
    Width  int    `json:"w"`
    Height int    `json:"h"`
    Resize string `json:"resize"`
}

type Sizes struct {
    Large  Size `json:"large"`
    Medium Size `json:"medium"`
    Small  Size `json:"small"`
    Thumb  Size `json:"thumb"`
}

type Medium struct {
    Id             int64     `json:"id"`
    IdStr          string    `json:"id_str"`
    Type           string    `json:"type"`
    MediaUrl       string    `json:"media_url"`
    SecureMediaUrl string    `json:"media_url_https"`
    Url            string    `json:"url"`
    DisplayUrl     string    `json:"display_url"`
    ExpandedUrl    *string   `json:"expanded_url"`
    Sizes          Sizes     `json:"sizes"`
    Indices        IndexPair `json:"indices"`
}

type Mention struct {
    Id         int64     `json:"id"`
    IdStr      string    `json:"id_str"`
    ScreenName string    `json:"screen_name"`
    FullName   string    `json:"full_name"`
    Indices    IndexPair `json:"indices"`
}

type Url struct {
    Url         string    `json:"url"`
    DisplayUrl  string    `json:"display_url"`
    ExpandedUrl *string   `json:"expanded_url"`
    Indices     IndexPair `json:"indices"`
}

type Entities struct {
    Hashtags []Hashtag `json:"hashtags"`
    Media    []Medium  `json:"media"`
    Mentions []Mention `json:"user_mentions"`
    Urls     []Url     `json:"urls"`
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
    Coordinates *Point `json:"coordinates"`

    // Time when this Tweet was created.
    CreatedAt Time `json:"created_at"`

    // Entities which have been parsed out of the text of the Tweet.
    Entities Entities `json:"entities"`

    // Perspectival. Indicates whether this Tweet has been favorited by the authenticating user.
    Favorited *bool `json:"favorited"`

    // If the represented Tweet is a reply, this field will contain the screen name of the original Tweet's author.
    InReplyToScreenName *string `json:"in_reply_to_screen_name"`

    // If the represented Tweet is a reply, this field will contain the integer representation of the original Tweet's ID.
    InReplyToStatusId *int64 `json:"in_reply_to_status_id"`

    // If the represented Tweet is a reply, this field will contain the string representation of the original Tweet's ID.
    InReplyToStatusIdStr *string `json:"in_reply_to_status_id_str"`

    // If the represented Tweet is a reply, this field will contain the integer representation of the original Tweet's author ID.
    InReplyToUserId *int64 `json:"in_reply_to_user_id"`

    // If the represented Tweet is a reply, this field will contain the string representation of the original Tweet's author ID.
    InReplyToUserIdStr *string `json:"in_reply_to_user_id_str"`

    // When present, indicates a BCP 47 language identifier corresponding to the machine-detected language of the Tweet text, or “und” if no language could be detected. 
    Lang string `json:"lang"`

    // When present, indicates that the tweet is associated (but not necessarily originating from) a Place.
    Place *Place `json:"place"`

    // This field only surfaces when a tweet contains a link. The meaning of the field doesn't pertain to the tweet content itself, but instead it is an indicator that the URL contained in the tweet may contain content or media identified as sensitive content.
    PossiblySensitive *bool `json:"possibly_sensitive"`

    // Number of times this Tweet has been retweeted. This field is no longer capped at 99 and will not turn into a String for "100+"
    RetweetCount int `json:"retweet_count"`

    // Perspectival. Indicates whether this Tweet has been retweeted by the authenticating user.
    Retweeted bool `json:"retweeted"`

    // If Retweet the original Tweet can be found here.
    RetweetedStatus *Tweet `json:"retweeted_status"`

    // Utility used to post the Tweet, as an HTML-formatted string. Tweets from the Twitter website have a source value of web.
    Source string `json:"source"`

    // The actual UTF-8 text of the status update.
    Text string `json:"text"`

    // Indicates whether the value of the text parameter was truncated, for example, as a result of a retweet exceeding the 140 character Tweet length. Truncated text will end in ellipsis, like this ...
    Truncated bool `json:"truncated"`

    // When present, indicates a textual representation of the two-letter country codes this content is withheld from. See New Withheld Content Fields in API Responses.
    WithheldInCountries string `json:"withheld_in_countries"`

    // When present, indicates whether the content being withheld is the "status" or a "user." See New Withheld Content Fields in API Responses.
    WithheldScope string `json:"withheld_scope"`

    // The user who posted this Tweet.
    User User `json:"user"`
}
