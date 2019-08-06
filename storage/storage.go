package storage

//go:generate retool do mockgen -destination=mocks/storage.go -package=mocks github.com/jloom6/phishql/storage Interface

import (
	"context"

	"github.com/jloom6/phishql/structs"
)

type Interface interface {
	GetShows(ctx context.Context, req structs.GetShowsRequest) ([]structs.Show, error)
	GetArtists(ctx context.Context, req structs.GetArtistsRequest) ([]structs.Artist, error)
	GetSongs(ctx context.Context, req structs.GetSongsRequest) ([]structs.Song, error)
}
