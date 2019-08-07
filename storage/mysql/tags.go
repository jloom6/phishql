package mysql

import (
	"context"

	"github.com/jloom6/phishql/internal/db"
	"github.com/jloom6/phishql/structs"
)

const (
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

// GetTags gets the tags from a MySQL database
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
