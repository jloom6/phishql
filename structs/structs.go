package structs

import "time"

type Artist struct {
	ID int `json:"id"`
	Name string `json:"name"`
}

type Venue struct {
	ID int `json:"id"`
	Name string `json:"name"`
	City string `json:"city"`
	State string `json:"state"`
}

type Tour struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Description string `json:"description,omitempty"`
}

type Song struct {
	ID int `json:"id"`
	Name string `json:"name"`
}

type Tag struct {
	ID int `json:"id"`
	Text string `json:"text"`
}

type SetSong struct {
	Song Song `json:"song"`
	Tag *Tag `json:"tag,omitempty"`
	Transition string `json:"transition,omitempty"`
}

type Set struct {
	ID int `json:"id"`
	Label string `json:"label"`
	Songs []SetSong `json:"songs"`
}

type Show struct {
	ID int `json:"id"`
	Date time.Time `json:"date"`
	Artist Artist `json:"artist"`
	Venue Venue `json:"venue"`
	Tour *Tour `json:"tour,omitempty"`
	Notes string `json:"notes,omitempty"`
	Soundcheck string `json:"soundcheck,omitempty"`
	Sets []Set `json:"sets,omitempty"`
}

type GetShowsRequest struct {

}
