package mapper

//go:generate retool do mockgen -destination=mocks/mapper.go -package=mocks github.com/jloom6/phishql/mapper Interface

import (
	"github.com/golang/protobuf/ptypes"
	phishqlpb "github.com/jloom6/phishql/.gen/proto/jloom6/phishql"
	"github.com/jloom6/phishql/structs"
)

type Interface interface {
	ProtoToGetShowsRequest(p *phishqlpb.GetShowsRequest) structs.GetShowsRequest
	ShowsToProto(shows []structs.Show) ([]*phishqlpb.Show, error)
}

type Mapper struct {}

func New() *Mapper {
	return &Mapper{}
}

func (m *Mapper) ProtoToGetShowsRequest(p *phishqlpb.GetShowsRequest) structs.GetShowsRequest {
	return structs.GetShowsRequest{
		Year: int(p.Year),
		Month: int(p.Month),
		Day: int(p.Day),
		DayOfWeek: int(p.DayOfWeek),
		City: p.City,
		Country: p.Country,
		State: p.State,
	}
}

func (m *Mapper) ShowsToProto(shows []structs.Show) ([]*phishqlpb.Show, error) {
	ps := make([]*phishqlpb.Show, 0, len(shows))

	for _, s := range shows {
		p, err := showToProto(s)
		if err != nil {
			return nil, err
		}

		ps = append(ps, p)
	}

	return ps, nil
}

func showToProto(s structs.Show) (*phishqlpb.Show, error) {
	t, err := ptypes.TimestampProto(s.Date)
	if err != nil {
		return nil, err
	}

	return &phishqlpb.Show{
		Id: int32(s.ID),
		Date: t,
		Artist: artistToProto(s.Artist),
		Venue: venueToProto(s.Venue),
		Tour: tourToProto(s.Tour),
		Notes: s.Notes,
		Soundcheck: s.Soundcheck,
		Sets: setsToProto(s.Sets),
	}, nil
}

func artistToProto(a structs.Artist) *phishqlpb.Artist {
	return &phishqlpb.Artist{
		Id: int32(a.ID),
		Name: a.Name,
	}
}

func venueToProto(v structs.Venue) *phishqlpb.Venue {
	return &phishqlpb.Venue{
		Id: int32(v.ID),
		Name: v.Name,
		City: v.City,
		State: v.State,
		Country: v.Country,
	}
}

func tourToProto(t *structs.Tour) *phishqlpb.Tour {
	if t == nil {
		return nil
	}

	return &phishqlpb.Tour{
		Id: int32(t.ID),
		Name: t.Name,
		Description: t.Description,
	}
}

func setsToProto(sets []structs.Set) []*phishqlpb.Set {
	ps := make([]*phishqlpb.Set, 0, len(sets))

	for _, s := range sets {
		ps = append(ps, setToProto(s))
	}

	return ps
}

func setToProto(s structs.Set) *phishqlpb.Set {
	return &phishqlpb.Set{
		Id: int32(s.ID),
		Label: s.Label,
		Songs: setSongsToProto(s.Songs),
	}
}

func setSongsToProto(songs []structs.SetSong) []*phishqlpb.SetSong {
	ps := make([]*phishqlpb.SetSong, 0, len(songs))

	for _, s := range songs {
		ps = append(ps, setSongToProto(s))
	}

	return ps
}

func setSongToProto(s structs.SetSong) *phishqlpb.SetSong {
	return &phishqlpb.SetSong{
		Song: songToProto(s.Song),
		Tag: tagToProto(s.Tag),
		Transition: s.Transition,
	}
}

func songToProto(s structs.Song) *phishqlpb.Song {
	return &phishqlpb.Song{
		Id: int32(s.ID),
		Name: s.Name,
	}
}

func tagToProto(t *structs.Tag) *phishqlpb.Tag {
	if t == nil {
		return nil
	}

	return &phishqlpb.Tag{
		Id: int32(t.ID),
		Text: t.Text,
	}
}
