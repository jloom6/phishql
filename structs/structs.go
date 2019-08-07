package structs

import (
	"time"
)

// Artist is an artist
type Artist struct {
	ID   int
	Name string
}

// Venue is a venue
type Venue struct {
	ID      int
	Name    string
	City    string
	State   string
	Country string
}

// Tour is a tour
type Tour struct {
	ID          int
	Name        string
	Description string
}

// Song is a song
type Song struct {
	ID   int
	Name string
}

// Tag is a tag
type Tag struct {
	ID   int
	Text string
}

// SetSong is the specific instance of a song being played at a show
type SetSong struct {
	Song       Song
	Tag        *Tag
	Transition string
}

// Set is a set
type Set struct {
	ID    int
	Label string
	Songs []SetSong
}

// Show is a show
type Show struct {
	ID         int
	Date       time.Time
	Artist     Artist
	Venue      Venue
	Tour       *Tour
	Notes      string
	Soundcheck string
	Sets       []Set
}
