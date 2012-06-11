package twitterstream

type User struct {
    // Indicates that the user has an account with "contributor mode" enabled, allowing for Tweets issued by the user to be co-authored by another account. Rarely true.
    ContributorsEnabled bool `json:"contributors_enabled"`

    // The UTC datetime that the user account was created on Twitter.
    CreatedAt Time `json:"created_at"`

    // When true, indicates that the user has not altered the theme or background of their user profile.
    DefaultProfile bool `json:"default_profile"`

    // When true, indicates that the user has not uploaded their own avatar and a default egg avatar is used instead.
    DefaultProfileImage bool `json:"default_profile_image"`

    // The user-defined UTF-8 string describing their account.
    Description *string `json:"description"`

    // The number of tweets this user has favorited in the account's lifetime. British spelling used in the field name for historical reasons.
    FavouritesCount int `json:"favourites_count"`

    // Perspectival. When true, indicates that the authenticating user has issued a follow request to this protected user account.
    FollowRequestSent *bool `json:"follow_request_sent"`

    // Perspectival. Deprecated. When true, indicates that the authenticating user is following this user. Some false negatives are possible when set to "false," but these false negatives are increasingly being represented as "null" instead. See Discussion.
    Following *bool `json:"following"`

    // The number of followers this account currently has. Under certain conditions of duress, this field will temporarily indicate "0."
    FollowersCount int `json:"followers_count"`

    // The number of users this account is following (AKA their "followings"). Under certain conditions of duress, this field will temporarily indicate "0."
    FriendsCount int `json:"friends_count"`

    // When true, indicates that the user has enabled the possibility of geotagging their Tweets. This field must be true for the current user to attach geographic data when using POST statuses/update.
    GeoEnabled bool `json:"geo_enabled"`

    // The integer representation of the unique identifier for this User. This number is greater than 53 bits and some programming languages may have difficulty/silent defects in interpreting it. Using a signed 64 bit integer for storing this identifier is safe. Use id_str for fetching the identifier to stay on the safe side.
    Id  int64 `json:"id"`

    // The string representation of the unique identifier for this Tweet. Implementations should use this rather than the large integer in id.
    IdStr string `json:"id_str"`

    // When true, indicates that the user is a participant in Twitter's translator community.
    IsTranslator bool `json:"is_translator"`

    // The ISO 639-1 two-letter character code for the user's self-declared user interface language. May or may not have anything to do with the content of their Tweets.
    Language string `json:"lang"`

    // The number of public lists that this user is a member of.
    ListedCount int `json:"listed_count"`

    // The user-defined location for this account's profile. Not necessarily a location nor parseable. This field will occasionally be fuzzily interpreted by the Search service.
    Location *string `json:"location"`

    // The name of the user, as they've defined it. Not necessarily a person's name. Typically capped at 20 characters, but subject to change.
    Name string `json:"name"`

    // The hexadecimal color chosen by the user for their background.
    ProfileBackgroundColor *string `json:"profile_background_color"`

    // A HTTP-based URL pointing to the background image the user has uploaded for their profile.
    ProfileBackgroundImageUrl *string `json:"profile_background_image_url"`

    // A HTTPS-based URL pointing to the background image the user has uploaded for their profile.
    ProfileBackgroundImageUrlHttps *string `json:"profile_background_image_url_https"`

    // When true, indicates that the user's profile_background_image_url should be tiled when displayed.
    ProfileBackgroundTile bool `json:"profile_background_tile"`

    // A HTTP-based URL pointing to the user's avatar image.
    ProfileImageUrl string `json:"profile_image_url"`

    // A HTTPS-based URL pointing to the user's avatar image.
    ProfileImageUrlHttps string `json:"profile_image_url_https"`

    // The hexadecimal color the user has chosen to display links with in their Twitter UI.
    ProfileLinkColor string `json:"profile_link_color"`

    // The hexadecimal color the user has chosen to display sidebar borders with in their Twitter UI.
    ProfileSidebarBorderColor string `json:"profile_sidebar_border_color"`

    // The hexadecimal color the user has chosen to display sidebar backgrounds with in their Twitter UI.
    ProfileSidebarFillColor string `json:"profile_sidebar_fill_color"`

    // The hexadecimal color the user has chosen to display text with in their Twitter UI.
    ProfileTextColor string `json:"profile_text_color"`

    // When true, indicates the user wants their uploaded background image to be used.
    ProfileUseBackgroundImage bool `json:"profile_use_background_image"`

    // When true, indicates that this user has chosen to protect their Tweets.
    Protected bool `json:"protected"`

    // The screen name, handle, or alias that this user identifies themselves with. screen_names are unique but subject to change. Use id_str as a user identifier whenever possible. Typically a maximum of 15 characters long, but some historical accounts may exist with longer names.
    ScreenName string `json:"screen_name"`

    // Indicates that the user would like to see media inline.
    ShowAllInlineMedia bool `json:"show_all_inline_media"`

    // If possible, the user's most recent tweet or retweet. In some circumstances, this data cannot be provided and this field will be omitted, null, or empty. Perspectival attributes within tweets embedded within users cannot always be relied upon.
    Status *Tweet `json:"status"`

    // The number of tweets (including retweets) issued by the user.
    StatusesCount int `json:"statuses_count"`

    // A string describing the Time Zone this user declares themselves within.
    TimeZone *string `json:"time_zone"`

    // A URL provided by the user in association with their profile.
    Url *string `json:"url"`

    // The offset from GMT/UTC in seconds.
    UtcOffset *int `json:"utc_offset"`

    // When true, indicates that the user has a verified account.
    Verified bool `json:"verified"`

    // When present, indicates a textual representation of the two-letter country codes this user is withheld from.
    WithheldInCountries string `json:"withheld_in_countries"`

    // When present, indicates whether the content being withheld is the "status" or a "user." See New Withheld Content Fields in API Responses.
    WithheldScope string `json:"withheld_scope"`
}
