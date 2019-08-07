package structs

// GetShowsRequest is a request to get shows
type GetShowsRequest struct {
	Condition Condition
}

// Condition allows you to compose base conditions
type Condition struct {
	Base BaseCondition
	Ands []Condition
	Ors  []Condition
}

// BaseCondition is the base conditions for querying
type BaseCondition struct {
	Year      int
	Month     int
	Day       int
	DayOfWeek int
	City      string
	State     string
	Country   string
	Song      string
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
