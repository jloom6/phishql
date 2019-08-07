package mysql

import (
	"context"
	"fmt"
	"strings"

	"github.com/jloom6/phishql/internal/db"
	"github.com/jloom6/phishql/structs"
)

const (
	baseWhereCondition = `
CASE WHEN ? = 0 THEN 1=1 ELSE YEAR(shows.date) = ? END AND
CASE WHEN ? = 0 THEN 1=1 ELSE MONTH(shows.date) = ? END AND
CASE WHEN ? = 0 THEN 1=1 ELSE DAY(shows.date) = ? END AND
CASE WHEN ? = 0 THEN 1=1 ELSE DAYOFWEEK(shows.date) = ? END AND
CASE WHEN ? = '' THEN 1=1 ELSE venues.city = ? END AND
CASE WHEN ? = '' THEN 1=1 ELSE venues.state = ? END AND
CASE WHEN ? = '' THEN 1=1 ELSE venues.country = ? END AND
CASE WHEN ? = '' THEN 1=1 ELSE
	? = ANY (
		SELECT
			songs.name
		FROM
			sets
				INNER JOIN set_songs ON sets.id = set_songs.set_id
					INNER JOIN songs ON set_songs.song_id = songs.id
		WHERE
			sets.show_id = shows.id
	)
END
`
	getShowsQuery = `
SELECT
	shows.id show_id,
	shows.date show_date,
	artists.id artist_id,
	artists.name artist_name,
	venues.id venue_id,
	venues.name venue_name,
	venues.city venue_city,
	venues.state venue_state,
	venues.country venue_country,
	COALESCE(tours.id, 0) tour_id,
	COALESCE(tours.name, '') tour_name,
	COALESCE(tours.desc, '') tour_description,
	COALESCE(shows.notes, '') show_notes,
	COALESCE(shows.soundcheck, '') show_soundcheck
FROM
	shows
		INNER JOIN artists ON shows.artist_id = artists.id
		INNER JOIN venues ON shows.venue_id = venues.id
		LEFT JOIN tours ON shows.tour_id = tours.id
WHERE
	%s
`
	getSetsQuery = `
SELECT
	sets.show_id show_id,
	sets.order set_order,
	sets.id set_id,
	sets.label set_label
FROM
	sets
		INNER JOIN shows ON sets.show_id = shows.id
			INNER JOIN venues on shows.venue_id = venues.id
WHERE
	%s
`
	getSetSongsQuery = `
SELECT
	sets.show_id show_id,
	sets.order set_order,
	songs.id song_id,
	songs.name song_name,
	COALESCE(tags.id, 0) tag_id,
	COALESCE(tags.text, '') tag_text,
	COALESCE(set_songs.transition, '') song_transition
FROM
	set_songs
		INNER JOIN sets ON set_songs.set_id = sets.id
			INNER JOIN shows ON sets.show_id = shows.id
				INNER JOIN venues on shows.venue_id = venues.id
		INNER JOIN songs ON set_songs.song_id = songs.id
		LEFT JOIN tags ON set_songs.tag_id = tags.id
WHERE
	%s
ORDER BY
	sets.show_id,
	sets.order,
	set_songs.order
`
	getArtistsQuery = `
SELECT
	id,
	name
FROM
	artists
ORDER BY
	name
`
	getSongsQuery = `
SELECT
	id,
	name
FROM
	songs
ORDER BY
	name
`
	getTagsQuery = `
SELECT
	id,
	text
FROM
	tags
ORDER BY
	text
`
)

type Store struct {
	db db.Interface
}

type Params struct {
	DB db.Interface
}

func New(p Params) *Store {
	return &Store{db: p.DB}
}

func (s *Store) GetShows(ctx context.Context, req structs.GetShowsRequest) ([]structs.Show, error) {
	shows, err := s.getShows(ctx, req)
	if err != nil {
		return nil, err
	}

	sets, err := s.getSets(ctx, req)
	if err != nil {
		return nil, err
	}

	songs, err := s.getSongs(ctx, req)
	if err != nil {
		return nil, err
	}

	return hydrateShows(shows, sets, songs), nil
}

func (s *Store) getShows(ctx context.Context, req structs.GetShowsRequest) (map[int]structs.Show, error) {
	query, args := makeQueryAndArgs(getShowsQuery, req.Condition)

	rows, err := s.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	defer closeRows(rows)

	return makeShows(rows)
}

func makeQueryAndArgs(baseQuery string, c structs.Condition) (string, []interface{}) {
	whereClause, args := makeClauseAndArgs(c)

	return fmt.Sprintf(baseQuery, whereClause), args
}

func makeClauseAndArgs(c structs.Condition) (string, []interface{}) {
	if c.Ands != nil {
		return makeAndClauseAndArgs(c)
	}

	if c.Ors != nil {
		return makeOrClauseAndArgs(c)
	}

	return makeBaseClauseAndArgs(c)
}

func makeAndClauseAndArgs(c structs.Condition) (string, []interface{}) {
	return makeClausesAndArgs(c.Ands, " AND ")
}

func makeClausesAndArgs(cs []structs.Condition, joiner string) (string, []interface{}) {
	var whereClauses []string
	var allArgs []interface{}

	for _, orCondition := range cs {
		whereClause, args := makeClauseAndArgs(orCondition)

		whereClauses = append(whereClauses, fmt.Sprintf("(%s)", whereClause))
		allArgs = append(allArgs, args...)
	}

	return strings.Join(whereClauses, joiner), allArgs
}

func makeOrClauseAndArgs(c structs.Condition) (string, []interface{}) {
	return makeClausesAndArgs(c.Ors, " OR ")
}

func makeBaseClauseAndArgs(c structs.Condition) (string, []interface{}) {
	return baseWhereCondition, makeBaseWhereArgs(c.Base)
}

func makeBaseWhereArgs(bc structs.BaseCondition) []interface{} {
	return []interface{}{
		bc.Year, bc.Year,
		bc.Month, bc.Month,
		bc.Day, bc.Day,
		bc.DayOfWeek, bc.DayOfWeek,
		bc.City, bc.City,
		bc.State, bc.State,
		bc.Country, bc.Country,
		bc.Song, bc.Song,
	}
}

func closeRows(rows db.Rows) {
	_ = rows.Close()
}

func makeShows(rows db.Rows) (map[int]structs.Show, error) {
	shows := make(map[int]structs.Show, 0)

	for rows.Next() {
		show, err := makeShow(rows)
		if err != nil {
			return nil, err
		}

		shows[show.ID] = show
	}

	return shows, nil
}

func makeShow(row db.Rows) (structs.Show, error) {
	var show structs.Show
	var tour structs.Tour

	err := row.Scan(&show.ID, &show.Date, &show.Artist.ID, &show.Artist.Name, &show.Venue.ID,
		&show.Venue.Name, &show.Venue.City, &show.Venue.State, &show.Venue.Country, &tour.ID,
		&tour.Name, &tour.Description, &show.Notes, &show.Soundcheck)
	if err != nil {
		return structs.Show{}, err
	}

	if tour != (structs.Tour{}) {
		show.Tour = &tour
	}

	return show, nil
}

func (s *Store) getSets(ctx context.Context, req structs.GetShowsRequest) (map[int]map[int]structs.Set, error) {
	query, args := makeQueryAndArgs(getSetsQuery, req.Condition)

	rows, err := s.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	defer closeRows(rows)

	return makeSets(rows)
}

func makeSets(rows db.Rows) (map[int]map[int]structs.Set, error) {
	sets := map[int]map[int]structs.Set{}

	for rows.Next() {
		var showID, setOrder, setID int
		var setLabel string

		if err := rows.Scan(&showID, &setOrder, &setID, &setLabel); err != nil {
			return nil, err
		}

		if _, ok := sets[showID]; !ok {
			sets[showID] = map[int]structs.Set{}
		}

		sets[showID][setOrder] = structs.Set{ID: setID, Label: setLabel}
	}

	return sets, nil
}

func (s *Store) getSongs(ctx context.Context, req structs.GetShowsRequest) (map[int]map[int][]structs.SetSong, error) {
	query, args := makeQueryAndArgs(getSetSongsQuery, req.Condition)

	rows, err := s.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	defer closeRows(rows)

	return makeSetSongs(rows)
}

func makeSetSongs(rows db.Rows) (map[int]map[int][]structs.SetSong, error) {
	songs := map[int]map[int][]structs.SetSong{}

	for rows.Next() {
		var showID, setOrder, songID int
		var songName, songTransition string
		var tag structs.Tag

		err := rows.Scan(&showID, &setOrder, &songID, &songName, &tag.ID, &tag.Text, &songTransition)
		if err != nil {
			return nil, err
		}

		if _, ok := songs[showID]; !ok {
			songs[showID] = map[int][]structs.SetSong{}
		}

		var tagP *structs.Tag
		if tag != (structs.Tag{}) {
			tagP = &tag
		}

		songs[showID][setOrder] = append(songs[showID][setOrder], structs.SetSong{
			Song: structs.Song{
				ID:   songID,
				Name: songName,
			},
			Tag:        tagP,
			Transition: songTransition,
		})
	}

	return songs, nil
}

func hydrateShows(shows map[int]structs.Show, sets map[int]map[int]structs.Set, songs map[int]map[int][]structs.SetSong) []structs.Show {
	hydratedShows := make([]structs.Show, 0, len(shows))

	for _, show := range shows {
		hydratedShows = append(hydratedShows, hydrateShow(show, sets, songs))
	}

	return hydratedShows
}

func hydrateShow(show structs.Show, sets map[int]map[int]structs.Set, songs map[int]map[int][]structs.SetSong) structs.Show {
	idx := 0
	for {
		set, hasMore := hydrateSet(sets[show.ID], songs[show.ID], idx)
		if !hasMore {
			break
		}

		idx++

		show.Sets = append(show.Sets, set)
	}

	return show
}

func hydrateSet(sets map[int]structs.Set, songs map[int][]structs.SetSong, idx int) (structs.Set, bool) {
	set, ok := sets[idx]
	if !ok {
		return structs.Set{}, false
	}

	if songs == nil {
		return structs.Set{}, true
	}

	if _, ok := songs[idx]; !ok {
		return structs.Set{}, true
	}

	set.Songs = songs[idx]

	return set, true
}

func (s *Store) GetArtists(ctx context.Context, _ structs.GetArtistsRequest) ([]structs.Artist, error) {
	rows, err := s.db.QueryContext(ctx, getArtistsQuery)
	if err != nil {
		return nil, err
	}

	defer closeRows(rows)

	return makeArtists(rows)
}

func makeArtists(rows db.Rows) ([]structs.Artist, error) {
	as := make([]structs.Artist, 0)

	for rows.Next() {
		a, err := makeArtist(rows)
		if err != nil {
			return nil, err
		}

		as = append(as, a)
	}

	return as, nil
}

func makeArtist(row db.Rows) (structs.Artist, error) {
	var a structs.Artist

	err := row.Scan(&a.ID, &a.Name)
	if err != nil {
		return structs.Artist{}, err
	}

	return a, nil
}

func (s *Store) GetSongs(ctx context.Context, _ structs.GetSongsRequest) ([]structs.Song, error) {
	rows, err := s.db.QueryContext(ctx, getSongsQuery)
	if err != nil {
		return nil, err
	}

	defer closeRows(rows)

	return makeSongs(rows)
}

func makeSongs(rows db.Rows) ([]structs.Song, error) {
	ss := make([]structs.Song, 0)

	for rows.Next() {
		s, err := makeSong(rows)
		if err != nil {
			return nil, err
		}

		ss = append(ss, s)
	}

	return ss, nil
}

func makeSong(row db.Rows) (structs.Song, error) {
	var s structs.Song

	err := row.Scan(&s.ID, &s.Name)
	if err != nil {
		return structs.Song{}, err
	}

	return s, nil
}

func (s *Store) GetTags(ctx context.Context, _ structs.GetTagsRequest) ([]structs.Tag, error) {
	rows, err := s.db.QueryContext(ctx, getTagsQuery)
	if err != nil {
		return nil, err
	}

	defer closeRows(rows)

	return makeTags(rows)
}

func makeTags(rows db.Rows) ([]structs.Tag, error) {
	ts := make([]structs.Tag, 0)

	for rows.Next() {
		t, err := makeTag(rows)
		if err != nil {
			return nil, err
		}

		ts = append(ts, t)
	}

	return ts, nil
}

func makeTag(row db.Rows) (structs.Tag, error) {
	var t structs.Tag

	err := row.Scan(&t.ID, &t.Text)
	if err != nil {
		return structs.Tag{}, err
	}

	return t, nil
}
