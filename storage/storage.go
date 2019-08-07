package storage

//go:generate retool do mockgen -destination=mocks/storage.go -package=mocks github.com/jloom6/phishql/storage Interface

import (
	"context"

	"github.com/jloom6/phishql/structs"
)

// Interface contains the storage logic
type Interface interface {
	GetShows(ctx context.Context, req structs.GetShowsRequest) ([]structs.Show, error)
	GetArtists(ctx context.Context, req structs.GetArtistsRequest) ([]structs.Artist, error)
	GetSongs(ctx context.Context, req structs.GetSongsRequest) ([]structs.Song, error)
	GetTags(ctx context.Context, req structs.GetTagsRequest) ([]structs.Tag, error)
	GetTours(ctx context.Context, req structs.GetToursRequest) ([]structs.Tour, error)
	GetVenues(ctx context.Context, req structs.GetVenuesRequest) ([]structs.Venue, error)
}
