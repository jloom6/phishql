package mysql

import (
	"context"

	"github.com/jloom6/phishql/internal/db"
	"github.com/jloom6/phishql/structs"
)

const (
	getVenuesQuery = `
SELECT
	id,
	name,
	city,
	state,
	country
FROM
	venues
ORDER BY
	name
`
)

// GetVenues gets the venues from a MySQL database
func (s *Store) GetVenues(ctx context.Context, _ structs.GetVenuesRequest) ([]structs.Venue, error) {
	rows, err := s.db.QueryContext(ctx, getVenuesQuery)
	if err != nil {
		return nil, err
	}

	defer closeRows(rows)

	return makeVenues(rows)
}

func makeVenues(rows db.Rows) ([]structs.Venue, error) {
	vs := make([]structs.Venue, 0)

	for rows.Next() {
		v, err := makeVenue(rows)
		if err != nil {
			return nil, err
		}

		vs = append(vs, v)
	}

	return vs, nil
}

func makeVenue(row db.Rows) (structs.Venue, error) {
	var v structs.Venue

	err := row.Scan(&v.ID, &v.Name, &v.City, &v.State, &v.Country)
	if err != nil {
		return structs.Venue{}, err
	}

	return v, nil
}
