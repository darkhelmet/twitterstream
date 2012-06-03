package twitterstream

type Place struct {
    // Contains a hash of variant information about the place.
    Attributes map[string]interface{} `json:"attributes"`

    // A bounding box of coordinates which encloses this place.
    BoundingBox Box `json:"bounding_box"`

    // Name of the country containing this place.
    Country string `json:"country"`

    // Shortened country code representing the country containing this place.
    CountryCode string `json:"country_code"`

    // Full human-readable representation of the place's name.
    FullName string `json:"full_name"`

    // ID representing this place. Note that this is represented as a string, not an integer.
    Id  string `json:"id"`

    // Short human-readable representation of the place's name.
    Name string `json:"name"`

    // The type of location represented by this place.
    Type string `json:"place_type"`

    // URL representing the location of additional place metadata for this place.
    Url string `json:"url"`
}
