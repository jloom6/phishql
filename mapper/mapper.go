package mapper

//go:generate retool do mockgen -destination=mocks/mapper.go -package=mocks github.com/jloom6/phishql/mapper Interface

import (
	"errors"

	"github.com/golang/protobuf/ptypes"
	phishqlpb "github.com/jloom6/phishql/.gen/proto/jloom6/phishql"
	"github.com/jloom6/phishql/structs"
)

// Interface wraps the mapping functions so which makes testing easier
type Interface interface {
	ProtoToGetShowsRequest(p *phishqlpb.GetShowsRequest) structs.GetShowsRequest
	ProtoToGetArtistsRequest(p *phishqlpb.GetArtistsRequest) structs.GetArtistsRequest
	ProtoToGetSongsRequest(p *phishqlpb.GetSongsRequest) structs.GetSongsRequest
	ProtoToGetTagsRequest(p *phishqlpb.GetTagsRequest) structs.GetTagsRequest
	ProtoToGetToursRequest(p *phishqlpb.GetToursRequest) structs.GetToursRequest
	ProtoToGetVenuesRequest(p *phishqlpb.GetVenuesRequest) structs.GetVenuesRequest
	ProtoToArtists(ps []*phishqlpb.Artist) []structs.Artist
	ProtoToShows(ps []*phishqlpb.Show) ([]structs.Show, error)
	ProtoToSongs(ps []*phishqlpb.Song) []structs.Song
	ProtoToTags(ps []*phishqlpb.Tag) []structs.Tag
	ProtoToTours(ps []*phishqlpb.Tour) []structs.Tour
	ProtoToVenues(ps []*phishqlpb.Venue) []structs.Venue
	ShowsToProto(ss []structs.Show) ([]*phishqlpb.Show, error)
	ArtistsToProto(as []structs.Artist) []*phishqlpb.Artist
	SongsToProto(ss []structs.Song) []*phishqlpb.Song
	TagsToProto(ts []structs.Tag) []*phishqlpb.Tag
	ToursToProto(ts []structs.Tour) []*phishqlpb.Tour
	VenuesToProto(vs []structs.Venue) []*phishqlpb.Venue
	GraphQLShowsToProto(args map[string]interface{}) (*phishqlpb.GetShowsRequest, error)
}

// Mapper implements the Interface to wrap mapping functions
type Mapper struct{}

// New constructs a new Mapper
func New() *Mapper {
	return &Mapper{}
}

// ProtoToGetShowsRequest maps a proto to struct
func (m *Mapper) ProtoToGetShowsRequest(p *phishqlpb.GetShowsRequest) structs.GetShowsRequest {
	return structs.GetShowsRequest{
		Condition: protoToCondition(p.Condition),
	}
}

func protoToCondition(p *phishqlpb.Condition) structs.Condition {
	return structs.Condition{
		Base: protoToBaseCondition(p.GetBase()),
		And:  protoToConditions(p.GetAnd()),
		Or:   protoToConditions(p.GetOr()),
	}
}

func protoToBaseCondition(p *phishqlpb.BaseCondition) structs.BaseCondition {
	if p == nil {
		return structs.BaseCondition{}
	}

	return structs.BaseCondition{
		Year:      int(p.Year),
		Month:     int(p.Month),
		Day:       int(p.Day),
		DayOfWeek: int(p.DayOfWeek),
		City:      p.City,
		Country:   p.Country,
		State:     p.State,
		Song:      p.Song,
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

// ShowsToProto maps shows to proto
func (m *Mapper) ShowsToProto(ss []structs.Show) ([]*phishqlpb.Show, error) {
	ps := make([]*phishqlpb.Show, 0, len(ss))

	for _, s := range ss {
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
		Id:         int32(s.ID),
		Date:       t,
		Artist:     artistToProto(s.Artist),
		Venue:      venueToProto(s.Venue),
		Tour:       tourToProto(s.Tour),
		Notes:      s.Notes,
		Soundcheck: s.Soundcheck,
		Sets:       setsToProto(s.Sets),
	}, nil
}

func artistToProto(a structs.Artist) *phishqlpb.Artist {
	return &phishqlpb.Artist{
		Id:   int32(a.ID),
		Name: a.Name,
	}
}

func venueToProto(v structs.Venue) *phishqlpb.Venue {
	return &phishqlpb.Venue{
		Id:      int32(v.ID),
		Name:    v.Name,
		City:    v.City,
		State:   v.State,
		Country: v.Country,
	}
}

func tourToProto(t *structs.Tour) *phishqlpb.Tour {
	if t == nil {
		return nil
	}

	return &phishqlpb.Tour{
		Id:          int32(t.ID),
		Name:        t.Name,
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
		Id:    int32(s.ID),
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
		Song:       songToProto(s.Song),
		Tag:        tagToProto(s.Tag),
		Transition: s.Transition,
	}
}

func songToProto(s structs.Song) *phishqlpb.Song {
	return &phishqlpb.Song{
		Id:   int32(s.ID),
		Name: s.Name,
	}
}

func tagToProto(t *structs.Tag) *phishqlpb.Tag {
	if t == nil {
		return nil
	}

	return &phishqlpb.Tag{
		Id:   int32(t.ID),
		Text: t.Text,
	}
}

// ProtoToGetArtistsRequest maps a proto to struct
func (m *Mapper) ProtoToGetArtistsRequest(p *phishqlpb.GetArtistsRequest) structs.GetArtistsRequest {
	return structs.GetArtistsRequest{}
}

// ArtistsToProto maps artists to proto
func (m *Mapper) ArtistsToProto(as []structs.Artist) []*phishqlpb.Artist {
	ps := make([]*phishqlpb.Artist, 0, len(as))

	for _, a := range as {
		ps = append(ps, artistToProto(a))
	}

	return ps
}

// ProtoToGetSongsRequest maps a proto to struct
func (m *Mapper) ProtoToGetSongsRequest(p *phishqlpb.GetSongsRequest) structs.GetSongsRequest {
	return structs.GetSongsRequest{}
}

// SongsToProto maps songs to proto
func (m *Mapper) SongsToProto(ss []structs.Song) []*phishqlpb.Song {
	ps := make([]*phishqlpb.Song, 0, len(ss))

	for _, s := range ss {
		ps = append(ps, songToProto(s))
	}

	return ps
}

// ProtoToGetTagsRequest maps a proto to struct
func (m *Mapper) ProtoToGetTagsRequest(p *phishqlpb.GetTagsRequest) structs.GetTagsRequest {
	return structs.GetTagsRequest{}
}

// TagsToProto maps tags to proto
func (m *Mapper) TagsToProto(ts []structs.Tag) []*phishqlpb.Tag {
	ps := make([]*phishqlpb.Tag, 0, len(ts))

	for _, t := range ts {
		ps = append(ps, tagToProto(&t))
	}

	return ps
}

// ProtoToGetToursRequest maps a proto to struct
func (m *Mapper) ProtoToGetToursRequest(p *phishqlpb.GetToursRequest) structs.GetToursRequest {
	return structs.GetToursRequest{}
}

// ToursToProto maps tours to proto
func (m *Mapper) ToursToProto(ts []structs.Tour) []*phishqlpb.Tour {
	ps := make([]*phishqlpb.Tour, 0, len(ts))

	for _, t := range ts {
		ps = append(ps, tourToProto(&t))
	}

	return ps
}

// ProtoToGetVenuesRequest maps a proto to struct
func (m *Mapper) ProtoToGetVenuesRequest(p *phishqlpb.GetVenuesRequest) structs.GetVenuesRequest {
	return structs.GetVenuesRequest{}
}

// VenuesToProto maps venues to proto
func (m *Mapper) VenuesToProto(vs []structs.Venue) []*phishqlpb.Venue {
	ps := make([]*phishqlpb.Venue, 0, len(vs))

	for _, v := range vs {
		ps = append(ps, venueToProto(v))
	}

	return ps
}

// ProtoToArtists maps a proto to struct
func (m *Mapper) ProtoToArtists(ps []*phishqlpb.Artist) []structs.Artist {
	as := make([]structs.Artist, 0, len(ps))

	for _, p := range ps {
		as = append(as, protoToArtist(p))
	}

	return as
}

func protoToArtist(p *phishqlpb.Artist) structs.Artist {
	if p == nil {
		return structs.Artist{}
	}

	return structs.Artist{
		ID:   int(p.Id),
		Name: p.Name,
	}
}

// ProtoToShows maps a proto to struct
func (m *Mapper) ProtoToShows(ps []*phishqlpb.Show) ([]structs.Show, error) {
	ss := make([]structs.Show, 0, len(ps))

	for _, p := range ps {
		s, err := protoToShow(p)
		if err != nil {
			return nil, err
		}

		ss = append(ss, s)
	}

	return ss, nil
}

func protoToShow(p *phishqlpb.Show) (structs.Show, error) {
	t, err := ptypes.Timestamp(p.Date)
	if err != nil {
		return structs.Show{}, err
	}

	return structs.Show{
		ID:         int(p.Id),
		Date:       t,
		Artist:     protoToArtist(p.Artist),
		Venue:      protoToVenue(p.Venue),
		Tour:       protoToTour(p.Tour),
		Notes:      p.Notes,
		Soundcheck: p.Soundcheck,
		Sets:       protoToSets(p.Sets),
	}, nil
}

func protoToSets(ps []*phishqlpb.Set) []structs.Set {
	ss := make([]structs.Set, 0, len(ps))

	for _, p := range ps {
		ss = append(ss, protoToSet(p))
	}

	return ss
}

func protoToSet(p *phishqlpb.Set) structs.Set {
	if p == nil {
		return structs.Set{}
	}

	return structs.Set{
		ID:    int(p.Id),
		Label: p.Label,
		Songs: protoToSetSongs(p.Songs),
	}
}

func protoToSetSongs(ps []*phishqlpb.SetSong) []structs.SetSong {
	ss := make([]structs.SetSong, 0, len(ps))

	for _, p := range ps {
		ss = append(ss, protoToSetSong(p))
	}

	return ss
}

func protoToSetSong(p *phishqlpb.SetSong) structs.SetSong {
	if p == nil {
		return structs.SetSong{}
	}

	return structs.SetSong{
		Song:       protoToSong(p.Song),
		Tag:        protoToTag(p.Tag),
		Transition: p.Transition,
	}
}

// ProtoToSongs maps a proto to struct
func (m *Mapper) ProtoToSongs(ps []*phishqlpb.Song) []structs.Song {
	ss := make([]structs.Song, 0, len(ps))

	for _, p := range ps {
		ss = append(ss, protoToSong(p))
	}

	return ss
}

func protoToSong(p *phishqlpb.Song) structs.Song {
	if p == nil {
		return structs.Song{}
	}

	return structs.Song{
		ID:   int(p.Id),
		Name: p.Name,
	}
}

// ProtoToTags maps a proto to struct
func (m *Mapper) ProtoToTags(ps []*phishqlpb.Tag) []structs.Tag {
	ts := make([]structs.Tag, 0, len(ps))

	for _, p := range ps {
		if t := protoToTag(p); t != nil {
			ts = append(ts, *t)
		}
	}

	return ts
}

func protoToTag(p *phishqlpb.Tag) *structs.Tag {
	if p == nil {
		return nil
	}

	return &structs.Tag{
		ID:   int(p.Id),
		Text: p.Text,
	}
}

// ProtoToVenues maps a proto to struct
func (m *Mapper) ProtoToTours(ps []*phishqlpb.Tour) []structs.Tour {
	ts := make([]structs.Tour, 0, len(ps))

	for _, p := range ps {
		if t := protoToTour(p); t != nil {
			ts = append(ts, *t)
		}
	}

	return ts
}

func protoToTour(p *phishqlpb.Tour) *structs.Tour {
	if p == nil {
		return nil
	}

	return &structs.Tour{
		ID:   int(p.Id),
		Name: p.Name,
	}
}

// ProtoToVenues maps a proto to struct
func (m *Mapper) ProtoToVenues(ps []*phishqlpb.Venue) []structs.Venue {
	vs := make([]structs.Venue, 0, len(ps))

	for _, p := range ps {
		vs = append(vs, protoToVenue(p))
	}

	return vs
}

func protoToVenue(p *phishqlpb.Venue) structs.Venue {
	if p == nil {
		return structs.Venue{}
	}

	return structs.Venue{
		ID:      int(p.Id),
		Name:    p.Name,
		City:    p.City,
		State:   p.State,
		Country: p.Country,
	}
}

// GraphQLShowsToProto maps a GraphQL shows request to proto
func (m *Mapper) GraphQLShowsToProto(args map[string]interface{}) (*phishqlpb.GetShowsRequest, error) {
	c, err := makeCondition(args["condition"])
	if err != nil {
		return nil, err
	}

	return &phishqlpb.GetShowsRequest{Condition:c}, nil
}

func makeCondition(arg interface{}) (*phishqlpb.Condition, error) {
	if arg == nil {
		return nil, errors.New("condition is nil")
	}

	args, ok := arg.(map[string]interface{})
	if !ok {
		return nil, errors.New("condition must be a map")
	}

	and, ok := args["and"]
	if ok {
		return makeAndCondition(and)
	}

	or, ok := args["or"]
	if ok {
		return makeOrCondition(or)
	}

	base, ok := args["base"]
	if ok {
		return makeBaseCondition(base)
	}

	return &phishqlpb.Condition{}, nil
}

func makeAndCondition(arg interface{}) (*phishqlpb.Condition, error) {
	cs, err := makeConditions(arg)
	if err != nil {
		return nil, err
	}

	return &phishqlpb.Condition{
		Condition: &phishqlpb.Condition_And{
			And: &phishqlpb.Conditions{
				Conditions: cs,
			},
		},
	}, nil
}

func makeConditions(arg interface{}) ([]*phishqlpb.Condition, error) {
	args, ok := arg.([]interface{})
	if !ok {
		return nil, errors.New("and condition must be a slice")
	}

	cs := make([]*phishqlpb.Condition, 0, len(args))

	for _, arg := range args {
		c, err := makeCondition(arg)
		if err != nil {
			return nil, err
		}

		cs = append(cs, c)
	}

	return cs, nil
}

func makeOrCondition(arg interface{}) (*phishqlpb.Condition, error) {
	cs, err := makeConditions(arg)
	if err != nil {
		return nil, err
	}

	return &phishqlpb.Condition{
		Condition: &phishqlpb.Condition_Or{
			Or: &phishqlpb.Conditions{
				Conditions: cs,
			},
		},
	}, nil
}

func makeBaseCondition(arg interface{}) (*phishqlpb.Condition, error) {
	args, ok := arg.(map[string]interface{})
	if !ok {
		return nil, errors.New("base condition must be a map")
	}

	bc := &phishqlpb.BaseCondition{}

	if arg, ok := args["year"]; ok {
		year, ok := arg.(int)
		if !ok {
			return nil, errors.New("year must be an int")
		}

		bc.Year = int32(year)
	}

	if arg, ok := args["month"]; ok {
		month, ok := arg.(int)
		if !ok {
			return nil, errors.New("month must be an int")
		}

		bc.Month = int32(month)
	}

	if arg, ok := args["day"]; ok {
		day, ok := arg.(int)
		if !ok {
			return nil, errors.New("day must be an int")
		}

		bc.Day = int32(day)
	}

	if arg, ok := args["dayOfWeek"]; ok {
		dayOfWeek, ok := arg.(int)
		if !ok {
			return nil, errors.New("dayOfWeek must be an int")
		}

		bc.DayOfWeek = int32(dayOfWeek)
	}

	if arg, ok := args["city"]; ok {
		bc.City, ok = arg.(string)
		if !ok {
			return nil, errors.New("city must be a string")
		}
	}

	if arg, ok := args["state"]; ok {
		bc.State, ok = arg.(string)
		if !ok {
			return nil, errors.New("state must be a string")
		}
	}

	if arg, ok := args["country"]; ok {
		bc.Country, ok = arg.(string)
		if !ok {
			return nil, errors.New("country must be a string")
		}
	}

	if arg, ok := args["song"]; ok {
		bc.Song, ok = arg.(string)
		if !ok {
			return nil, errors.New("song must be a string")
		}
	}

	return &phishqlpb.Condition{
		Condition: &phishqlpb.Condition_Base{
			Base: bc,
		},
	}, nil
}
