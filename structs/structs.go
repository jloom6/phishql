package structs

import "time"

type Artist struct {
	ID int
	Name string
}

type Venue struct {
	ID int
	Name string
	City string
	State string
}

type Tour struct {
	ID int
	Name string
	Description string
}

type Song struct {
	ID int
	Name string
}

type Tag struct {
	ID int
	Text string
}

type SetSong struct {
	Song Song
	Tag *Tag
	Transition string
}

type Set struct {
	ID int
	Label string
	Songs []SetSong
}

type Show struct {
	ID int
	Date time.Time
	Artist Artist
	Venue Venue
	Tour *Tour
	Notes string
	Soundcheck string
	Sets []Set
}

type GetShowsRequest struct {

}
