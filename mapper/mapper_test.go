package mapper

import (
	"errors"
	"testing"
	"time"

	"github.com/golang/protobuf/ptypes"
	phishqlpb "github.com/jloom6/phishql/.gen/proto/jloom6/phishql"
	"github.com/jloom6/phishql/structs"
	"github.com/stretchr/testify/assert"
)

func TestMapper_ProtoToGetShowsRequest(t *testing.T) {
	req := &phishqlpb.GetShowsRequest{
		Condition: &phishqlpb.Condition{
			Condition: &phishqlpb.Condition_And{
				And: &phishqlpb.Conditions{
					Conditions: []*phishqlpb.Condition{
						{
							Condition: &phishqlpb.Condition_Base{
								Base: &phishqlpb.BaseCondition{
									Year: 1994,
									Month: 10,
									Day: 31,
									DayOfWeek: 2,
									City: "Glens Falls",
									State: "NY",
									Country: "USA",
									Song: "Reba",
								},
							},
						},
					},
				},
			},
		},

	}

	tests := []struct{
		name string
		req *phishqlpb.GetShowsRequest
		ret structs.GetShowsRequest
	}{
		{
			name: "non-nil request",
			req: req,
			ret: structs.GetShowsRequest{
				Condition: structs.Condition{
					Ands: []structs.Condition{
						{
							Base: structs.BaseCondition{
								Year: int(req.Condition.GetAnd().Conditions[0].GetBase().Year),
								Month: int(req.Condition.GetAnd().Conditions[0].GetBase().Month),
								Day: int(req.Condition.GetAnd().Conditions[0].GetBase().Day),
								DayOfWeek: int(req.Condition.GetAnd().Conditions[0].GetBase().DayOfWeek),
								City: req.Condition.GetAnd().Conditions[0].GetBase().City,
								State: req.Condition.GetAnd().Conditions[0].GetBase().State,
								Country: req.Condition.GetAnd().Conditions[0].GetBase().Country,
								Song: req.Condition.GetAnd().Conditions[0].GetBase().Song,
							},
						},
					},
				},
			},
		},
	}

	m := New()

	for _, test := range tests {
		t.Run(test.name, func(_t *testing.T) {
			assert.Equal(_t, test.ret, m.ProtoToGetShowsRequest(test.req))
		})
	}
}

func TestMapper_ShowsToProto(t *testing.T) {
	invalidDateShow := structs.Show{
		Date: time.Time{}.Add(-1 * time.Second),
	}

	now := time.Now()

	show := structs.Show{
		ID: 1,
		Date: now,
		Artist: structs.Artist{
			ID: -7,
			Name: "Bob Weaver",
		},
		Venue: structs.Venue{
			ID: 60,
			Name: "Madison Square Garden",
			City: "New York",
			State: "NY",
			Country: "USA",
		},
		Tour: &structs.Tour{
			ID: 2017,
			Name: "The Baker's Dozen",
			Description: "Phish owns MSG for 13 nights",
		},
		Notes: "The show was 🔥",
		Soundcheck: "Jennifer Dances",
		Sets: []structs.Set{
			{
				ID: 1,
				Label: "SET 1",
				Songs: []structs.SetSong{
					{
						Song: structs.Song{
							ID: 555,
							Name: "555",
						},
						Tag: &structs.Tag{
							ID: 10,
							Text: "With Vacuum Solo",
						},
						Transition: "->",
					},
				},
			},
		},
	}

	pNow, _ := ptypes.TimestampProto(now)

	p := &phishqlpb.Show{
		Id: int32(show.ID),
		Date: pNow,
		Artist: &phishqlpb.Artist{
			Id: int32(show.Artist.ID),
			Name: show.Artist.Name,
		},
		Venue: &phishqlpb.Venue{
			Id: int32(show.Venue.ID),
			Name: show.Venue.Name,
			City: show.Venue.City,
			State: show.Venue.State,
			Country: show.Venue.Country,
		},
		Tour: &phishqlpb.Tour{
			Id: int32(show.Tour.ID),
			Name: show.Tour.Name,
			Description: show.Tour.Description,
		},
		Notes: show.Notes,
		Soundcheck: show.Soundcheck,
		Sets: []*phishqlpb.Set{
			{
				Id: int32(show.Sets[0].ID),
				Label: show.Sets[0].Label,
				Songs: []*phishqlpb.SetSong{
					{
						Song: &phishqlpb.Song{
							Id: int32(show.Sets[0].Songs[0].Song.ID),
							Name: show.Sets[0].Songs[0].Song.Name,
						},
						Tag: &phishqlpb.Tag{
							Id: int32(show.Sets[0].Songs[0].Tag.ID),
							Text: show.Sets[0].Songs[0].Tag.Text,
						},
						Transition: show.Sets[0].Songs[0].Transition,
					},
				},
			},
		},
	}

	tests := []struct{
		name string
		shows []structs.Show
		ret []*phishqlpb.Show
		err error
	}{
		{
			name: "ptypes.TimestampProto error",
			shows: []structs.Show{invalidDateShow},
			err: errors.New("timestamp: seconds:-62135596801  before 0001-01-01"),
		},
		{
			name: "success",
			shows: []structs.Show{show},
			ret: []*phishqlpb.Show{p},
		},
	}

	m := New()

	for _, test := range tests {
		t.Run(test.name, func(_t *testing.T) {
			ret, err := m.ShowsToProto(test.shows)

			assert.Equal(_t, test.ret, ret)
			assert.Equal(_t, test.err, err)
		})
	}
}

func TestTourToProto(t *testing.T) {
	tour := &structs.Tour{
		ID: 1998,
		Name: "The Island Tour",
	}

	p := &phishqlpb.Tour{
		Id: int32(tour.ID),
		Name: tour.Name,
	}

	tests := []struct{
		name string
		tour *structs.Tour
		ret *phishqlpb.Tour
	}{
		{
			name: "nil tour",
		},
		{
			name: "non-nil tour",
			tour: tour,
			ret: p,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(_t *testing.T) {
			assert.Equal(_t, test.ret, tourToProto(test.tour))
		})
	}
}

func TestTagToProto(t *testing.T) {
	tag := &structs.Tag{
		ID: 1,
		Text: "FYF",
	}

	p := &phishqlpb.Tag{
		Id: int32(tag.ID),
		Text: tag.Text,
	}

	tests := []struct{
		name string
		tag *structs.Tag
		ret *phishqlpb.Tag
	}{
		{
			name: "nil tag",
		},
		{
			name: "non-nil tag",
			tag: tag,
			ret: p,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(_t *testing.T) {
			assert.Equal(_t, test.ret, tagToProto(test.tag))
		})
	}
}

func TestMapper_ProtoToGetArtistsRequest(t *testing.T) {
	tests := []struct{
		name string
		req *phishqlpb.GetArtistsRequest
		ret structs.GetArtistsRequest
	}{
		{
			name: "nil req",
		},
		{
			name: "non-nil req",
			req: &phishqlpb.GetArtistsRequest{},
		},
	}

	m := New()

	for _, test := range tests {
		t.Run(test.name, func(_t *testing.T) {
			assert.Equal(_t, test.ret, m.ProtoToGetArtistsRequest(test.req))
		})
	}
}

func TestMapper_ArtistsToProto(t *testing.T) {
	a := structs.Artist{
		ID:   9,
		Name: "Kasvot Växt",
	}

	p := &phishqlpb.Artist{
		Id:                   int32(a.ID),
		Name:                 a.Name,
	}
	
	tests := []struct{
		name string
		artists []structs.Artist
		ret []*phishqlpb.Artist
	}{
		{
			name: "success",
			artists: []structs.Artist{a},
			ret: []*phishqlpb.Artist{p},
		},
	}

	m := New()

	for _, test := range tests {
		t.Run(test.name, func(_t *testing.T) {
			assert.Equal(_t, test.ret, m.ArtistsToProto(test.artists))
		})
	}
}

func TestMapper_ProtoToGetSongsRequest(t *testing.T) {
	tests := []struct{
		name string
		req *phishqlpb.GetSongsRequest
		ret structs.GetSongsRequest
	}{
		{
			name: "nil req",
		},
		{
			name: "non-nil req",
			req: &phishqlpb.GetSongsRequest{},
		},
	}

	m := New()

	for _, test := range tests {
		t.Run(test.name, func(_t *testing.T) {
			assert.Equal(_t, test.ret, m.ProtoToGetSongsRequest(test.req))
		})
	}
}

func TestMapper_SongsToProto(t *testing.T) {
	s := structs.Song{
		ID:   555,
		Name: "555",
	}

	p := &phishqlpb.Song{
		Id:                   int32(s.ID),
		Name:                 s.Name,
	}

	tests := []struct{
		name string
		songs []structs.Song
		ret []*phishqlpb.Song
	}{
		{
			name: "success",
			songs: []structs.Song{s},
			ret: []*phishqlpb.Song{p},
		},
	}

	m := New()

	for _, test := range tests {
		t.Run(test.name, func(_t *testing.T) {
			assert.Equal(_t, test.ret, m.SongsToProto(test.songs))
		})
	}
}

func TestMapper_ProtoToGetTagsRequest(t *testing.T) {
	tests := []struct{
		name string
		req *phishqlpb.GetTagsRequest
		ret structs.GetTagsRequest
	}{
		{
			name: "nil req",
		},
		{
			name: "non-nil req",
			req: &phishqlpb.GetTagsRequest{},
		},
	}

	m := New()

	for _, test := range tests {
		t.Run(test.name, func(_t *testing.T) {
			assert.Equal(_t, test.ret, m.ProtoToGetTagsRequest(test.req))
		})
	}
}

func TestMapper_TagsToProto(t *testing.T) {
	tag := structs.Tag{
		ID:   1,
		Text: "Trey on marimba lumina.",
	}

	p := &phishqlpb.Tag{
		Id:                   int32(tag.ID),
		Text:                 tag.Text,
	}

	tests := []struct{
		name string
		tags []structs.Tag
		ret []*phishqlpb.Tag
	}{
		{
			name: "success",
			tags: []structs.Tag{tag},
			ret: []*phishqlpb.Tag{p},
		},
	}

	m := New()

	for _, test := range tests {
		t.Run(test.name, func(_t *testing.T) {
			assert.Equal(_t, test.ret, m.TagsToProto(test.tags))
		})
	}
}

func TestMapper_ProtoToGetToursRequest(t *testing.T) {
	tests := []struct{
		name string
		req *phishqlpb.GetToursRequest
		ret structs.GetToursRequest
	}{
		{
			name: "nil req",
		},
		{
			name: "non-nil req",
			req: &phishqlpb.GetToursRequest{},
		},
	}

	m := New()

	for _, test := range tests {
		t.Run(test.name, func(_t *testing.T) {
			assert.Equal(_t, test.ret, m.ProtoToGetToursRequest(test.req))
		})
	}
}

func TestMapper_ToursToProto(t *testing.T) {
	tour := structs.Tour{
		ID:   1,
		Name: "2015 Summer Tour",
		Description: "The end of set spelling at Dick's",
	}

	p := &phishqlpb.Tour{
		Id:                   int32(tour.ID),
		Name:                 tour.Name,
		Description: tour.Description,
	}

	tests := []struct{
		name string
		tours []structs.Tour
		ret []*phishqlpb.Tour
	}{
		{
			name: "success",
			tours: []structs.Tour{tour},
			ret: []*phishqlpb.Tour{p},
		},
	}

	m := New()

	for _, test := range tests {
		t.Run(test.name, func(_t *testing.T) {
			assert.Equal(_t, test.ret, m.ToursToProto(test.tours))
		})
	}
}
