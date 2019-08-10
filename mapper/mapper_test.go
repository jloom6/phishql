package mapper

import (
	"errors"
	"testing"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
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
									Year:      1994,
									Month:     10,
									Day:       31,
									DayOfWeek: 2,
									City:      "Glens Falls",
									State:     "NY",
									Country:   "USA",
									Song:      "Reba",
								},
							},
						},
					},
				},
			},
		},
	}

	tests := []struct {
		name string
		req  *phishqlpb.GetShowsRequest
		ret  structs.GetShowsRequest
	}{
		{
			name: "non-nil request",
			req:  req,
			ret: structs.GetShowsRequest{
				Condition: structs.Condition{
					And: []structs.Condition{
						{
							Base: structs.BaseCondition{
								Year:      int(req.Condition.GetAnd().Conditions[0].GetBase().Year),
								Month:     int(req.Condition.GetAnd().Conditions[0].GetBase().Month),
								Day:       int(req.Condition.GetAnd().Conditions[0].GetBase().Day),
								DayOfWeek: int(req.Condition.GetAnd().Conditions[0].GetBase().DayOfWeek),
								City:      req.Condition.GetAnd().Conditions[0].GetBase().City,
								State:     req.Condition.GetAnd().Conditions[0].GetBase().State,
								Country:   req.Condition.GetAnd().Conditions[0].GetBase().Country,
								Song:      req.Condition.GetAnd().Conditions[0].GetBase().Song,
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
		ID:   1,
		Date: now,
		Artist: structs.Artist{
			ID:   -7,
			Name: "Bob Weaver",
		},
		Venue: structs.Venue{
			ID:      60,
			Name:    "Madison Square Garden",
			City:    "New York",
			State:   "NY",
			Country: "USA",
		},
		Tour: &structs.Tour{
			ID:          2017,
			Name:        "The Baker's Dozen",
			Description: "Phish owns MSG for 13 nights",
		},
		Notes:      "The show was 🔥",
		Soundcheck: "Jennifer Dances",
		Sets: []structs.Set{
			{
				ID:    1,
				Label: "SET 1",
				Songs: []structs.SetSong{
					{
						Song: structs.Song{
							ID:   555,
							Name: "555",
						},
						Tag: &structs.Tag{
							ID:   10,
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
		Id:   int32(show.ID),
		Date: pNow,
		Artist: &phishqlpb.Artist{
			Id:   int32(show.Artist.ID),
			Name: show.Artist.Name,
		},
		Venue: &phishqlpb.Venue{
			Id:      int32(show.Venue.ID),
			Name:    show.Venue.Name,
			City:    show.Venue.City,
			State:   show.Venue.State,
			Country: show.Venue.Country,
		},
		Tour: &phishqlpb.Tour{
			Id:          int32(show.Tour.ID),
			Name:        show.Tour.Name,
			Description: show.Tour.Description,
		},
		Notes:      show.Notes,
		Soundcheck: show.Soundcheck,
		Sets: []*phishqlpb.Set{
			{
				Id:    int32(show.Sets[0].ID),
				Label: show.Sets[0].Label,
				Songs: []*phishqlpb.SetSong{
					{
						Song: &phishqlpb.Song{
							Id:   int32(show.Sets[0].Songs[0].Song.ID),
							Name: show.Sets[0].Songs[0].Song.Name,
						},
						Tag: &phishqlpb.Tag{
							Id:   int32(show.Sets[0].Songs[0].Tag.ID),
							Text: show.Sets[0].Songs[0].Tag.Text,
						},
						Transition: show.Sets[0].Songs[0].Transition,
					},
				},
			},
		},
	}

	tests := []struct {
		name  string
		shows []structs.Show
		ret   []*phishqlpb.Show
		err   error
	}{
		{
			name:  "ptypes.TimestampProto error",
			shows: []structs.Show{invalidDateShow},
			err:   errors.New("timestamp: seconds:-62135596801  before 0001-01-01"),
		},
		{
			name:  "success",
			shows: []structs.Show{show},
			ret:   []*phishqlpb.Show{p},
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
		ID:   1998,
		Name: "The Island Tour",
	}

	p := &phishqlpb.Tour{
		Id:   int32(tour.ID),
		Name: tour.Name,
	}

	tests := []struct {
		name string
		tour *structs.Tour
		ret  *phishqlpb.Tour
	}{
		{
			name: "nil tour",
		},
		{
			name: "non-nil tour",
			tour: tour,
			ret:  p,
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
		ID:   1,
		Text: "FYF",
	}

	p := &phishqlpb.Tag{
		Id:   int32(tag.ID),
		Text: tag.Text,
	}

	tests := []struct {
		name string
		tag  *structs.Tag
		ret  *phishqlpb.Tag
	}{
		{
			name: "nil tag",
		},
		{
			name: "non-nil tag",
			tag:  tag,
			ret:  p,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(_t *testing.T) {
			assert.Equal(_t, test.ret, tagToProto(test.tag))
		})
	}
}

func TestMapper_ProtoToGetArtistsRequest(t *testing.T) {
	tests := []struct {
		name string
		req  *phishqlpb.GetArtistsRequest
		ret  structs.GetArtistsRequest
	}{
		{
			name: "nil req",
		},
		{
			name: "non-nil req",
			req:  &phishqlpb.GetArtistsRequest{},
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
		Id:   int32(a.ID),
		Name: a.Name,
	}

	tests := []struct {
		name    string
		artists []structs.Artist
		ret     []*phishqlpb.Artist
	}{
		{
			name:    "success",
			artists: []structs.Artist{a},
			ret:     []*phishqlpb.Artist{p},
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
	tests := []struct {
		name string
		req  *phishqlpb.GetSongsRequest
		ret  structs.GetSongsRequest
	}{
		{
			name: "nil req",
		},
		{
			name: "non-nil req",
			req:  &phishqlpb.GetSongsRequest{},
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
		Id:   int32(s.ID),
		Name: s.Name,
	}

	tests := []struct {
		name  string
		songs []structs.Song
		ret   []*phishqlpb.Song
	}{
		{
			name:  "success",
			songs: []structs.Song{s},
			ret:   []*phishqlpb.Song{p},
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
	tests := []struct {
		name string
		req  *phishqlpb.GetTagsRequest
		ret  structs.GetTagsRequest
	}{
		{
			name: "nil req",
		},
		{
			name: "non-nil req",
			req:  &phishqlpb.GetTagsRequest{},
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
		Id:   int32(tag.ID),
		Text: tag.Text,
	}

	tests := []struct {
		name string
		tags []structs.Tag
		ret  []*phishqlpb.Tag
	}{
		{
			name: "success",
			tags: []structs.Tag{tag},
			ret:  []*phishqlpb.Tag{p},
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
	tests := []struct {
		name string
		req  *phishqlpb.GetToursRequest
		ret  structs.GetToursRequest
	}{
		{
			name: "nil req",
		},
		{
			name: "non-nil req",
			req:  &phishqlpb.GetToursRequest{},
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
		ID:          1,
		Name:        "2015 Summer Tour",
		Description: "The end of set spelling at Dick's",
	}

	p := &phishqlpb.Tour{
		Id:          int32(tour.ID),
		Name:        tour.Name,
		Description: tour.Description,
	}

	tests := []struct {
		name  string
		tours []structs.Tour
		ret   []*phishqlpb.Tour
	}{
		{
			name:  "success",
			tours: []structs.Tour{tour},
			ret:   []*phishqlpb.Tour{p},
		},
	}

	m := New()

	for _, test := range tests {
		t.Run(test.name, func(_t *testing.T) {
			assert.Equal(_t, test.ret, m.ToursToProto(test.tours))
		})
	}
}

func TestMapper_ProtoToGetVenuesRequest(t *testing.T) {
	tests := []struct {
		name string
		req  *phishqlpb.GetVenuesRequest
		ret  structs.GetVenuesRequest
	}{
		{
			name: "nil req",
		},
		{
			name: "non-nil req",
			req:  &phishqlpb.GetVenuesRequest{},
		},
	}

	m := New()

	for _, test := range tests {
		t.Run(test.name, func(_t *testing.T) {
			assert.Equal(_t, test.ret, m.ProtoToGetVenuesRequest(test.req))
		})
	}
}

func TestMapper_ProtoToArtists(t *testing.T) {
	a := structs.Artist{
		ID:   9,
		Name: "Kasvot Växt",
	}

	p := &phishqlpb.Artist{
		Id:   int32(a.ID),
		Name: a.Name,
	}

	tests := []struct {
		name    string
		proto []*phishqlpb.Artist
		ret     []structs.Artist
	}{
		{
			name: "nil artist",
			proto: []*phishqlpb.Artist{nil},
			ret: []structs.Artist{{}},
		},
		{
			name:    "success",
			proto: []*phishqlpb.Artist{p},
			ret:     []structs.Artist{a},
		},
	}

	m := New()

	for _, test := range tests {
		t.Run(test.name, func(_t *testing.T) {
			assert.Equal(_t, test.ret, m.ProtoToArtists(test.proto))
		})
	}
}

func TestMapper_ProtoToShows(t *testing.T) {
	s := structs.Show{
		ID: 2,
		Date: time.Date(2019, 3, 17, 0, 0, 0, 0, time.UTC),
		Artist: structs.Artist{
			ID:   3,
			Name: "Mike Gordon",
		},
		Venue: structs.Venue{
			ID:      2009,
			Name:    "Hampton Coliseum",
			City:    "Hampton",
			State:   "VA",
			Country: "USA",
		},
		Tour: &structs.Tour{
			ID:          1,
			Name:        "2015 Summer Tour",
			Description: "The end of set spelling at Dick's",
		},
		Notes: "these are notes",
		Soundcheck: "soundcheck",
		Sets: []structs.Set{
			{
				ID:    1,
				Label: "ENCORE 2",
				Songs: []structs.SetSong{
					{
						Song: structs.Song{
							ID:   555,
							Name: "555",
						},
						Tag: &structs.Tag{
							ID:   1,
							Text: "Trey on marimba lumina.",
						},
						Transition: "->",
					},
				},
			},
		},
	}

	pDate := &timestamp.Timestamp{Seconds: s.Date.Unix()}

	p := &phishqlpb.Show{
		Id: int32(s.ID),
		Date: pDate,
		Artist: &phishqlpb.Artist{
			Id:                   int32(s.Artist.ID),
			Name:                 s.Artist.Name,
		},
		Venue: &phishqlpb.Venue{
			Id:                   int32(s.Venue.ID),
			Name:                 s.Venue.Name,
			City:                 s.Venue.City,
			State:                s.Venue.State,
			Country:              s.Venue.Country,
		},
		Tour: &phishqlpb.Tour{
			Id:                   int32(s.Tour.ID),
			Name:                 s.Tour.Name,
			Description:          s.Tour.Description,
		},
		Notes: s.Notes,
		Soundcheck: s.Soundcheck,
		Sets: []*phishqlpb.Set{
			{
				Id:    int32(s.Sets[0].ID),
				Label: s.Sets[0].Label,
				Songs: []*phishqlpb.SetSong{
					{
						Song: &phishqlpb.Song{
							Id:   int32(s.Sets[0].Songs[0].Song.ID),
							Name: s.Sets[0].Songs[0].Song.Name,
						},
						Tag: &phishqlpb.Tag{
							Id:   int32(s.Sets[0].Songs[0].Tag.ID),
							Text: s.Sets[0].Songs[0].Tag.Text,
						},
						Transition: s.Sets[0].Songs[0].Transition,
					},
				},
			},
		},
	}

	tests := []struct {
		name  string
		proto []*phishqlpb.Show
		ret   []structs.Show
		err   error
	}{
		{
			name:  "nil show",
			proto: []*phishqlpb.Show{nil},
			ret:   []structs.Show{{}},
		},
		{
			name:  "protoToShow error",
			proto: []*phishqlpb.Show{{}},
			err: errors.New("timestamp: nil Timestamp"),
		},
		{
			name:  "success",
			proto: []*phishqlpb.Show{p},
			ret:   []structs.Show{s},
		},
	}

	m := New()

	for _, test := range tests {
		t.Run(test.name, func(_t *testing.T) {
			ret, err := m.ProtoToShows(test.proto)

			assert.Equal(_t, test.ret, ret)
			assert.Equal(_t, test.err, err)
		})
	}
}

func TestProtoToSet(t *testing.T) {
	s := structs.Set{
		ID: 1,
		Label: "ENCORE 2",
		Songs: []structs.SetSong{
			{
				Song: structs.Song{
					ID:   555,
					Name: "555",
				},
				Tag: &structs.Tag{
					ID:   1,
					Text: "Trey on marimba lumina.",
				},
				Transition: "->",
			},
		},
	}

	p := &phishqlpb.Set{
		Id:    int32(s.ID),
		Label: s.Label,
		Songs: []*phishqlpb.SetSong{
			{
				Song: &phishqlpb.Song{
					Id:   int32(s.Songs[0].Song.ID),
					Name: s.Songs[0].Song.Name,
				},
				Tag: &phishqlpb.Tag{
					Id:   int32(s.Songs[0].Tag.ID),
					Text: s.Songs[0].Tag.Text,
				},
				Transition: s.Songs[0].Transition,
			},
		},
	}

	tests := []struct {
		name  string
		proto *phishqlpb.Set
		ret   structs.Set
	}{
		{
			name:  "nil set",
			ret:   structs.Set{},
		},
		{
			name:  "success",
			proto: p,
			ret:   s,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(_t *testing.T) {
			assert.Equal(_t, test.ret, protoToSet(test.proto))
		})
	}
}

func TestProtoToSetSong(t *testing.T) {
	s := structs.SetSong{
		Song: structs.Song{
			ID:   555,
			Name: "555",
		},
		Tag: &structs.Tag{
			ID:   1,
			Text: "Trey on marimba lumina.",
		},
		Transition: "->",
	}

	p := &phishqlpb.SetSong{
		Song: &phishqlpb.Song{
			Id:                   int32(s.Song.ID),
			Name:                 s.Song.Name,
		},
		Tag: &phishqlpb.Tag{
			Id:                   int32(s.Tag.ID),
			Text:                 s.Tag.Text,
		},
		Transition: s.Transition,
	}

	tests := []struct {
		name  string
		proto *phishqlpb.SetSong
		ret   structs.SetSong
	}{
		{
			name:  "nil set song",
			ret:   structs.SetSong{},
		},
		{
			name:  "success",
			proto: p,
			ret:   s,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(_t *testing.T) {
			assert.Equal(_t, test.ret, protoToSetSong(test.proto))
		})
	}
}

func TestMapper_ProtoToSongs(t *testing.T) {
	s := structs.Song{
		ID:   555,
		Name: "555",
	}

	p := &phishqlpb.Song{
		Id:   int32(s.ID),
		Name: s.Name,
	}

	tests := []struct {
		name  string
		proto []*phishqlpb.Song
		ret   []structs.Song
	}{
		{
			name:  "nil song",
			proto: []*phishqlpb.Song{nil},
			ret:   []structs.Song{{}},
		},
		{
			name:  "success",
			proto: []*phishqlpb.Song{p},
			ret:   []structs.Song{s},
		},
	}

	m := New()

	for _, test := range tests {
		t.Run(test.name, func(_t *testing.T) {
			assert.Equal(_t, test.ret, m.ProtoToSongs(test.proto))
		})
	}
}

func TestMapper_ProtoToTags(t *testing.T) {
	tag := structs.Tag{
		ID:   1,
		Text: "Trey on marimba lumina.",
	}

	p := &phishqlpb.Tag{
		Id:   int32(tag.ID),
		Text: tag.Text,
	}

	tests := []struct {
		name string
		proto []*phishqlpb.Tag
		ret  []structs.Tag
	}{
		{
			name: "nil tag",
			proto: []*phishqlpb.Tag{nil},
			ret:  []structs.Tag{},
		},
		{
			name: "success",
			proto: []*phishqlpb.Tag{p},
			ret:  []structs.Tag{tag},
		},
	}

	m := New()

	for _, test := range tests {
		t.Run(test.name, func(_t *testing.T) {
			assert.Equal(_t, test.ret, m.ProtoToTags(test.proto))
		})
	}
}

func TestMapper_ProtoToTours(t *testing.T) {
	tour := structs.Tour{
		ID:          1,
		Name:        "2015 Summer Tour",
		Description: "The end of set spelling at Dick's",
	}

	p := &phishqlpb.Tour{
		Id:          int32(tour.ID),
		Name:        tour.Name,
		Description: tour.Description,
	}

	tests := []struct {
		name  string
		proto []*phishqlpb.Tour
		ret   []structs.Tour
	}{
		{
			name: "nil tour",
			proto: []*phishqlpb.Tour{nil},
			ret:   []structs.Tour{},
		},
		{
			name:  "success",
			proto: []*phishqlpb.Tour{p},
			ret:   []structs.Tour{tour},
		},
	}

	m := New()

	for _, test := range tests {
		t.Run(test.name, func(_t *testing.T) {
			assert.Equal(_t, test.ret, m.ProtoToTours(test.proto))
		})
	}
}

func TestMapper_VenuesToProto(t *testing.T) {
	v := structs.Venue{
		ID:      2009,
		Name:    "Hampton Coliseum",
		City:    "Hampton",
		State:   "VA",
		Country: "USA",
	}

	p := &phishqlpb.Venue{
		Id:      int32(v.ID),
		Name:    v.Name,
		City:    v.City,
		State:   v.State,
		Country: v.Country,
	}

	tests := []struct {
		name   string
		venues []structs.Venue
		ret    []*phishqlpb.Venue
	}{
		{
			name:   "success",
			venues: []structs.Venue{v},
			ret:    []*phishqlpb.Venue{p},
		},
	}

	m := New()

	for _, test := range tests {
		t.Run(test.name, func(_t *testing.T) {
			assert.Equal(_t, test.ret, m.VenuesToProto(test.venues))
		})
	}
}

func TestMapper_ProtoToVenues(t *testing.T) {
	v := structs.Venue{
		ID:      2009,
		Name:    "Hampton Coliseum",
		City:    "Hampton",
		State:   "VA",
		Country: "USA",
	}

	p := &phishqlpb.Venue{
		Id:      int32(v.ID),
		Name:    v.Name,
		City:    v.City,
		State:   v.State,
		Country: v.Country,
	}

	tests := []struct {
		name   string
		proto []*phishqlpb.Venue
		ret    []structs.Venue
	}{
		{
			name:   "nil venue",
			proto: []*phishqlpb.Venue{nil},
			ret:    []structs.Venue{{}},
		},
		{
			name:   "success",
			proto: []*phishqlpb.Venue{p},
			ret:    []structs.Venue{v},
		},
	}

	m := New()

	for _, test := range tests {
		t.Run(test.name, func(_t *testing.T) {
			assert.Equal(_t, test.ret, m.ProtoToVenues(test.proto))
		})
	}
}

func TestMapper_GraphQLShowsToProto(t *testing.T) {
	year := 1994

	req := &phishqlpb.GetShowsRequest{
		Condition: &phishqlpb.Condition{
			Condition: &phishqlpb.Condition_Base{
				Base: &phishqlpb.BaseCondition{
					Year: int32(year),
				},
			},
		},
	}

	tests := []struct {
		name string
		args  map[string]interface{}
		ret  *phishqlpb.GetShowsRequest
		err  error
	}{
		{
			name: "makeCondition error",
			err:  errors.New("condition is nil"),
		},
		{
			name: "success",
			args: map[string]interface{}{
				"condition": map[string]interface{}{
					"base": map[string]interface{}{
						"year": year,
					},
				},
			},
			ret: req,
		},
	}

	m := New()

	for _, test := range tests {
		t.Run(test.name, func(_t *testing.T) {
			ret, err := m.GraphQLShowsToProto(test.args)

			assert.Equal(_t, test.ret, ret)
			assert.Equal(_t, test.err, err)
		})
	}
}

func TestMakeCondition(t *testing.T) {
	year := 1994

	c := &phishqlpb.Condition{
		Condition: &phishqlpb.Condition_And{
			And: &phishqlpb.Conditions{
				Conditions: []*phishqlpb.Condition{
					{
						Condition: &phishqlpb.Condition_Or{
							Or: &phishqlpb.Conditions{
								Conditions: []*phishqlpb.Condition{
									{
										Condition: &phishqlpb.Condition_Base{
											Base: &phishqlpb.BaseCondition{
												Year: int32(year),
											},
										},
									},
									{
										Condition: &phishqlpb.Condition_Base{
											Base: &phishqlpb.BaseCondition{},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	tests := []struct {
		name string
		arg  interface{}
		ret  *phishqlpb.Condition
		err  error
	}{
		{
			name: "condition is nil",
			err:  errors.New("condition is nil"),
		},
		{
			name: "condition not a map",
			arg: 1,
			err: errors.New("condition must be a map"),
		},
		{
			name: "success",
			arg: map[string]interface{}{
				"and": []interface{}{
					map[string]interface{}{
						"or": []interface{}{
							map[string]interface{}{
								"base": map[string]interface{}{
									"year": year,
								},
							},
							map[string]interface{}{},
						},
					},
				},
			},
			ret: c,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(_t *testing.T) {
			ret, err := makeCondition(test.arg)

			assert.Equal(_t, test.ret, ret)
			assert.Equal(_t, test.err, err)
		})
	}
}

func TestMakeAndCondition(t *testing.T) {
	year := 1994

	c := &phishqlpb.Condition{
		Condition: &phishqlpb.Condition_And{
			And: &phishqlpb.Conditions{
				Conditions: []*phishqlpb.Condition{
					{
						Condition: &phishqlpb.Condition_Base{
							Base: &phishqlpb.BaseCondition{
								Year: int32(year),
							},
						},
					},
				},
			},
		},
	}

	tests := []struct {
		name string
		arg  interface{}
		ret  *phishqlpb.Condition
		err  error
	}{
		{
			name: "make conditions error",
			err:  errors.New("conditions must be a slice"),
		},
		{
			name: "success",
			arg: []interface{}{
				map[string]interface{}{
					"base": map[string]interface{}{
						"year": year,
					},
				},
			},
			ret: c,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(_t *testing.T) {
			ret, err := makeAndCondition(test.arg)

			assert.Equal(_t, test.ret, ret)
			assert.Equal(_t, test.err, err)
		})
	}
}

func TestMakeConditions(t *testing.T) {
	year := 1994

	c := []*phishqlpb.Condition{
		{
			Condition: &phishqlpb.Condition_Base{
				Base: &phishqlpb.BaseCondition{
					Year: int32(year),
				},
			},
		},
	}

	tests := []struct {
		name string
		arg  interface{}
		ret  []*phishqlpb.Condition
		err  error
	}{
		{
			name: "not a slice",
			err:  errors.New("conditions must be a slice"),
		},
		{
			name: "makeCondition error",
			arg: []interface{}{nil},
			err:  errors.New("condition is nil"),
		},
		{
			name: "success",
			arg: []interface{}{
				map[string]interface{}{
					"base": map[string]interface{}{
						"year": year,
					},
				},
			},
			ret: c,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(_t *testing.T) {
			ret, err := makeConditions(test.arg)

			assert.Equal(_t, test.ret, ret)
			assert.Equal(_t, test.err, err)
		})
	}
}

func TestMakeOrCondition(t *testing.T) {
	year := 1994

	c := &phishqlpb.Condition{
		Condition: &phishqlpb.Condition_Or{
			Or: &phishqlpb.Conditions{
				Conditions: []*phishqlpb.Condition{
					{
						Condition: &phishqlpb.Condition_Base{
							Base: &phishqlpb.BaseCondition{
								Year: int32(year),
							},
						},
					},
				},
			},
		},
	}

	tests := []struct {
		name string
		arg  interface{}
		ret  *phishqlpb.Condition
		err  error
	}{
		{
			name: "make conditions error",
			err:  errors.New("conditions must be a slice"),
		},
		{
			name: "success",
			arg: []interface{}{
				map[string]interface{}{
					"base": map[string]interface{}{
						"year": year,
					},
				},
			},
			ret: c,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(_t *testing.T) {
			ret, err := makeOrCondition(test.arg)

			assert.Equal(_t, test.ret, ret)
			assert.Equal(_t, test.err, err)
		})
	}
}

func TestMakeBaseCondition(t *testing.T) {
	bc := &phishqlpb.BaseCondition{
		Year: 2018,
		Month: 10,
		Day: 31,
		DayOfWeek: 4,
		City: "Las Vegas",
		State: "NV",
		Country: "USA",
		Song: "Say It To Me S.A.N.T.O.S.",
	}

	tests := []struct {
		name string
		arg  interface{}
		ret  *phishqlpb.Condition
		err  error
	}{
		{
			name: "not a map",
			err:  errors.New("base condition must be a map"),
		},
		{
			name: "year not an int",
			arg: map[string]interface{}{
				"year": "not an int",
			},
			err:  errors.New("year must be an int"),
		},
		{
			name: "month not an int",
			arg: map[string]interface{}{
				"month": "not an int",
			},
			err:  errors.New("month must be an int"),
		},
		{
			name: "day not an int",
			arg: map[string]interface{}{
				"day": "not an int",
			},
			err:  errors.New("day must be an int"),
		},
		{
			name: "dayOfWeek not an int",
			arg: map[string]interface{}{
				"dayOfWeek": "not an int",
			},
			err:  errors.New("dayOfWeek must be an int"),
		},
		{
			name: "city not a string",
			arg: map[string]interface{}{
				"city": 1,
			},
			err:  errors.New("city must be a string"),
		},
		{
			name: "state not a string",
			arg: map[string]interface{}{
				"state": 1,
			},
			err:  errors.New("state must be a string"),
		},
		{
			name: "country not a string",
			arg: map[string]interface{}{
				"country": 1,
			},
			err:  errors.New("country must be a string"),
		},
		{
			name: "song not a string",
			arg: map[string]interface{}{
				"song": 1,
			},
			err:  errors.New("song must be a string"),
		},
		{
			name: "success",
			arg: map[string]interface{}{
				"year": int(bc.Year),
				"month": int(bc.Month),
				"day": int(bc.Day),
				"dayOfWeek": int(bc.DayOfWeek),
				"city": bc.City,
				"state": bc.State,
				"country": bc.Country,
				"song": bc.Song,
			},
			ret: &phishqlpb.Condition{
				Condition: &phishqlpb.Condition_Base{
					Base: bc,
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(_t *testing.T) {
			ret, err := makeBaseCondition(test.arg)

			assert.Equal(_t, test.ret, ret)
			assert.Equal(_t, test.err, err)
		})
	}
}
