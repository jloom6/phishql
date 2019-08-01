package storage

import (
	"context"

	"github.com/jloom6/phishql/structs"
)

type Interface interface {
	GetShows(ctx context.Context, req structs.GetShowsRequest) ([]structs.Show, error)
}
