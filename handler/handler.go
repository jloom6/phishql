package handler

import (
	"context"
	"log"

	phishqlpb "github.com/jloom6/phishql/.gen/proto/jloom6/phishql"
	"github.com/jloom6/phishql/mapper"
	"github.com/jloom6/phishql/service"
)

type Handler struct {
	service service.Interface
	mapper mapper.Interface
}

type Params struct {
	Service service.Interface
	Mapper mapper.Interface
}

func New(p Params) *Handler {
	return &Handler{
		service: p.Service,
		mapper: p.Mapper,
	}
}

func (h *Handler) GetShows(ctx context.Context, req *phishqlpb.GetShowsRequest) (*phishqlpb.GetShowsResponse, error) {
	log.Printf("%v", req)

	shows, err := h.service.GetShows(ctx, h.mapper.ProtoToGetShowsRequest(req))
	if err != nil {
		return nil, err
	}

	ps, err := h.mapper.ShowsToProto(shows)
	if err != nil {
		return nil, err
	}

	return &phishqlpb.GetShowsResponse{Shows: ps}, nil
}

func (h *Handler) GetArtists(ctx context.Context, req *phishqlpb.GetArtistsRequest) (*phishqlpb.GetArtistsResponse, error) {
	log.Printf("%v", req)

	artists, err := h.service.GetArtists(ctx, h.mapper.ProtoToGetArtistsRequest(req))
	if err != nil {
		return nil, err
	}

	return &phishqlpb.GetArtistsResponse{
		Artists: h.mapper.ArtistsToProto(artists),
	}, nil
}

func (h *Handler) GetSongs(ctx context.Context, req *phishqlpb.GetSongsRequest) (*phishqlpb.GetSongsResponse, error) {
	log.Printf("%v", req)

	songs, err := h.service.GetSongs(ctx, h.mapper.ProtoToGetSongsRequest(req))
	if err != nil {
		return nil, err
	}

	return &phishqlpb.GetSongsResponse{
		Songs: h.mapper.SongsToProto(songs),
	}, nil
}

func (h *Handler) GetTags(ctx context.Context, req *phishqlpb.GetTagsRequest) (*phishqlpb.GetTagsResponse, error) {
	log.Printf("%v", req)

	tags, err := h.service.GetTags(ctx, h.mapper.ProtoToGetTagsRequest(req))
	if err != nil {
		return nil, err
	}

	return &phishqlpb.GetTagsResponse{
		Tags: h.mapper.TagsToProto(tags),
	}, nil
}

func (h *Handler) GetTours(ctx context.Context, req *phishqlpb.GetToursRequest) (*phishqlpb.GetToursResponse, error) {
	log.Printf("%v", req)

	tours, err := h.service.GetTours(ctx, h.mapper.ProtoToGetToursRequest(req))
	if err != nil {
		return nil, err
	}

	return &phishqlpb.GetToursResponse{
		Tours: h.mapper.ToursToProto(tours),
	}, nil
}
