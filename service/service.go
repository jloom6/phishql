package service

import (
	"context"
	"sort"

	"github.com/jloom6/phishql/storage"
	"github.com/jloom6/phishql/structs"
)

type Interface interface {
	GetShows(ctx context.Context, req structs.GetShowsRequest) ([]structs.Show, error)
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
