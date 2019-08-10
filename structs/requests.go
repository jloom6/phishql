package structs

// GetShowsRequest is a request to get shows
type GetShowsRequest struct {
	Condition Condition
}

// Condition allows you to compose base conditions
type Condition struct {
	Base BaseCondition `json:"base"`
	And  []Condition   `json:"and"`
	Or   []Condition   `json:"or"`
}

// BaseCondition is the base conditions for querying
type BaseCondition struct {
	Year      int    `json:"year"`
	Month     int    `json:"month"`
	Day       int    `json:"day"`
	DayOfWeek int    `json:"dayOfWeek"`
	City      string `json:"city"`
	State     string `json:"state"`
	Country   string `json:"country"`
	Song      string `json:"song"`
}

// GetArtistsRequest is a request to get artists
type GetArtistsRequest struct {
}

// GetSongsRequest is a request to get songs
type GetSongsRequest struct {
}

// GetTagsRequest is a request to get tags
type GetTagsRequest struct {
}

// GetToursRequest is a request to get tours
type GetToursRequest struct {
}

// GetVenuesRequest is a request to get venues
type GetVenuesRequest struct {
}
