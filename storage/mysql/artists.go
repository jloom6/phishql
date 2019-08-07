package mysql

import (
	"context"

	"github.com/jloom6/phishql/internal/db"
	"github.com/jloom6/phishql/structs"
)

const (
	getArtistsQuery = `
SELECT
	id,
	name
FROM
	artists
ORDER BY
	name
`
)

// GetArtists gets the artists from a MySQL database
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
