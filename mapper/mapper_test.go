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
	tests := []struct{
		name string
		req *phishqlpb.GetShowsRequest
		ret structs.GetShowsRequest
	}{
		{
			name: "nil request",
			ret: structs.GetShowsRequest{},
		},
		{
			name: "non-nil request",
			req: &phishqlpb.GetShowsRequest{},
			ret: structs.GetShowsRequest{},
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
