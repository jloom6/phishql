package service

//go:generate retool do mockgen -destination=mocks/service.go -package=mocks github.com/jloom6/phishql/service Interface

import (
	"context"
	"sort"

	"github.com/jloom6/phishql/storage"
	"github.com/jloom6/phishql/structs"
)

type Interface interface {
	GetShows(ctx context.Context, req structs.GetShowsRequest) ([]structs.Show, error)
	GetArtists(ctx context.Context, req structs.GetArtistsRequest) ([]structs.Artist, error)
	GetSongs(ctx context.Context, req structs.GetSongsRequest) ([]structs.Song, error)
	GetTags(ctx context.Context, req structs.GetTagsRequest) ([]structs.Tag, error)
	GetTours(ctx context.Context, req structs.GetToursRequest) ([]structs.Tour, error)
}

type Service struct {
	store storage.Interface
}

type Params struct {
	Store storage.Interface
}

func New(p Params) *Service {
	return &Service{store: p.Store}
}

func (s *Service) GetShows(ctx context.Context, req structs.GetShowsRequest) ([]structs.Show, error) {
	shows, err := s.store.GetShows(context.Background(), req)
	if err != nil {
		return nil, err
	}

	sort.Slice(shows, func(i, j int) bool { return shows[i].Date.Before(shows[j].Date) })

	return shows, nil
}

func (s *Service) GetArtists(ctx context.Context, req structs.GetArtistsRequest) ([]structs.Artist, error) {
	return s.store.GetArtists(ctx, req)
}

func (s *Service) GetSongs(ctx context.Context, req structs.GetSongsRequest) ([]structs.Song, error) {
	return s.store.GetSongs(ctx, req)
}

func (s *Service) GetTags(ctx context.Context, req structs.GetTagsRequest) ([]structs.Tag, error) {
	return s.store.GetTags(ctx, req)
}

func (s *Service) GetTours(ctx context.Context, req structs.GetToursRequest) ([]structs.Tour, error) {
	return s.store.GetTours(ctx, req)
}
