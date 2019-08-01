package mysql

import (
	"context"

	"github.com/jloom6/phishql/internal/db"
	"github.com/jloom6/phishql/structs"
)

const (
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
`
	getSetsQuery = `
SELECT
	sets.show_id show_id,
    sets.order set_order,
	sets.id set_id,
    sets.label set_label
FROM
	sets
`
	getSongsQuery = `
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
		INNER JOIN songs ON set_songs.song_id = songs.id
		LEFT JOIN tags ON set_songs.tag_id = tags.id
ORDER BY
	sets.show_id,
    sets.order,
    set_songs.order
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

func (s *Store) getShows(ctx context.Context, _ structs.GetShowsRequest) (map[int]structs.Show, error) {
	rows, err := s.db.QueryContext(ctx, getShowsQuery)
	if err != nil {
		return nil, err
	}

	defer closeRows(rows)

	return makeShows(rows)
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
		&show.Venue.Name, &show.Venue.City, &show.Venue.State, &tour.ID, &tour.Name,
		&tour.Description, &show.Notes, &show.Soundcheck)
	if err != nil {
		return structs.Show{}, err
	}

	if tour != (structs.Tour{}) {
		show.Tour = &tour
	}

	return show, nil
}

func (s *Store) getSets(ctx context.Context, _ structs.GetShowsRequest) (map[int]map[int]structs.Set, error) {
	rows, err := s.db.QueryContext(ctx, getSetsQuery)
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

func (s *Store) getSongs(ctx context.Context, _ structs.GetShowsRequest) (map[int]map[int][]structs.SetSong, error) {
	rows, err := s.db.QueryContext(ctx, getSongsQuery)
	if err != nil {
		return nil, err
	}

	defer closeRows(rows)

	return makeSongs(rows)
}

func makeSongs(rows db.Rows) (map[int]map[int][]structs.SetSong, error) {
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

		tagP := &tag
		if tag == (structs.Tag{}) {
			tagP = nil
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
