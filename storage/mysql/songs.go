package mysql

import (
	"context"

	"github.com/jloom6/phishql/internal/db"
	"github.com/jloom6/phishql/structs"
)

const (
	getSongsQuery = `
SELECT
	id,
	name
FROM
	songs
ORDER BY
	name
`
)

// GetSongs gets the songs from a MySQL database
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
