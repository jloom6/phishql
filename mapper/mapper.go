package mapper

//go:generate retool do mockgen -destination=mocks/mapper.go -package=mocks github.com/jloom6/phishql/mapper Interface

import (
	"github.com/golang/protobuf/ptypes"
	phishqlpb "github.com/jloom6/phishql/.gen/proto/jloom6/phishql"
	"github.com/jloom6/phishql/structs"
)

type Interface interface {
	ProtoToGetShowsRequest(p *phishqlpb.GetShowsRequest) structs.GetShowsRequest
	ProtoToGetArtistsRequest(p *phishqlpb.GetArtistsRequest) structs.GetArtistsRequest
	ProtoToGetSongsRequest(p *phishqlpb.GetSongsRequest) structs.GetSongsRequest
	ProtoToGetTagsRequest(p *phishqlpb.GetTagsRequest) structs.GetTagsRequest
	ShowsToProto(shows []structs.Show) ([]*phishqlpb.Show, error)
	ArtistsToProto(as []structs.Artist) []*phishqlpb.Artist
	SongsToProto(ss []structs.Song) []*phishqlpb.Song
	TagsToProto(ts []structs.Tag) []*phishqlpb.Tag
}

type Mapper struct {}

func New() *Mapper {
	return &Mapper{}
}

func (m *Mapper) ProtoToGetShowsRequest(p *phishqlpb.GetShowsRequest) structs.GetShowsRequest {
	return structs.GetShowsRequest{
		Condition: protoToCondition(p.Condition),
	}
}

func protoToCondition(p *phishqlpb.Condition) structs.Condition {
	return structs.Condition{
		Base: protoToBaseCondition(p.GetBase()),
		Ands: protoToConditions(p.GetAnd()),
		Ors: protoToConditions(p.GetOr()),
	}
}

func protoToBaseCondition(p *phishqlpb.BaseCondition) structs.BaseCondition {
	if p == nil {
		return structs.BaseCondition{}
	}

	return structs.BaseCondition{
		Year: int(p.Year),
		Month: int(p.Month),
		Day: int(p.Day),
		DayOfWeek: int(p.DayOfWeek),
		City: p.City,
		Country: p.Country,
		State: p.State,
		Song: p.Song,
	}
}

func protoToConditions(ps *phishqlpb.Conditions) []structs.Condition {
	if ps == nil {
		return nil
	}

	conditions := make([]structs.Condition, 0, len(ps.Conditions))

	for _, p := range ps.Conditions {
		conditions = append(conditions, protoToCondition(p))
	}

	return conditions
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

func (m *Mapper) ProtoToGetArtistsRequest(p *phishqlpb.GetArtistsRequest) structs.GetArtistsRequest {
	return structs.GetArtistsRequest{}
}

func (m *Mapper) ArtistsToProto(as []structs.Artist) []*phishqlpb.Artist {
	ps := make([]*phishqlpb.Artist, 0, len(as))

	for _, a := range as {
		ps = append(ps, artistToProto(a))
	}

	return ps
}

func (m *Mapper) ProtoToGetSongsRequest(p *phishqlpb.GetSongsRequest) structs.GetSongsRequest {
	return structs.GetSongsRequest{}
}

func (m *Mapper) SongsToProto(ss []structs.Song) []*phishqlpb.Song {
	ps := make([]*phishqlpb.Song, 0, len(ss))

	for _, s := range ss {
		ps = append(ps, songToProto(s))
	}

	return ps
}

func (m *Mapper) ProtoToGetTagsRequest(p *phishqlpb.GetTagsRequest) structs.GetTagsRequest {
	return structs.GetTagsRequest{}
}

func (m *Mapper) TagsToProto(ts []structs.Tag) []*phishqlpb.Tag {
	ps := make([]*phishqlpb.Tag, 0, len(ts))

	for _, t := range ts {
		ps = append(ps, tagToProto(&t))
	}

	return ps
}
