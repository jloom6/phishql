package mysql

import (
	"context"

	"github.com/jloom6/phishql/internal/db"
	"github.com/jloom6/phishql/structs"
)

const (
	// Go doesn't like backticks in string literals
	getToursQuery = `
SELECT
	id,
	name,
` +
		"`desc`\n" + `
FROM
	tours
ORDER BY
	name
`
)

// GetTours gets the tours from a MySQL database
func (s *Store) GetTours(ctx context.Context, _ structs.GetToursRequest) ([]structs.Tour, error) {
	rows, err := s.db.QueryContext(ctx, getToursQuery)
	if err != nil {
		return nil, err
	}

	defer closeRows(rows)

	return makeTours(rows)
}

func makeTours(rows db.Rows) ([]structs.Tour, error) {
	ts := make([]structs.Tour, 0)

	for rows.Next() {
		t, err := makeTour(rows)
		if err != nil {
			return nil, err
		}

		ts = append(ts, t)
	}

	return ts, nil
}

func makeTour(row db.Rows) (structs.Tour, error) {
	var t structs.Tour

	err := row.Scan(&t.ID, &t.Name, &t.Description)
	if err != nil {
		return structs.Tour{}, err
	}

	return t, nil
}
