package structs

import (
	"time"
)

// Artist is an artist
type Artist struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Venue is a venue
type Venue struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	City    string `json:"city"`
	State   string `json:"state"`
	Country string `json:"country"`
}

// Tour is a tour
type Tour struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// Song is a song
type Song struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Tag is a tag
type Tag struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
}

// SetSong is the specific instance of a song being played at a show
type SetSong struct {
	Song       Song   `json:"song"`
	Tag        *Tag   `json:"tag"`
	Transition string `json:"transition"`
}

// Set is a set
type Set struct {
	ID    int       `json:"id"`
	Label string    `json:"label"`
	Songs []SetSong `json:"songs"`
}

// Show is a show
type Show struct {
	ID         int       `json:"id"`
	Date       time.Time `json:"date"`
	Artist     Artist    `json:"artist"`
	Venue      Venue     `json:"venue"`
	Tour       *Tour     `json:"tour"`
	Notes      string    `json:"notes"`
	Soundcheck string    `json:"soundcheck"`
	Sets       []Set     `json:"sets"`
}
