package handler

import (
	"context"
	"log"

	phishqlpb "github.com/jloom6/phishql/.gen/proto/jloom6/phishql"
	"github.com/jloom6/phishql/mapper"
	"github.com/jloom6/phishql/service"
)

// Handler implements the PhishQL gRPC server interface
type Handler struct {
	service service.Interface
	mapper  mapper.Interface
}

// Params contains the parameters needed to construct a new Handler
type Params struct {
	Service service.Interface
	Mapper  mapper.Interface
}

// New constructs a mew Handler
func New(p Params) *Handler {
	return &Handler{
		service: p.Service,
		mapper:  p.Mapper,
	}
}

// GetShows gets the shows that meet the conditions from the request
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

// GetArtists gets the artists that meet the conditions from the request
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

// GetSongs gets the songs that meet the conditions from the request
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

// GetTags gets the tags that meet the conditions from the request
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

// GetTours gets the tours that meet the conditions from the request
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

// GetVenues gets the venues that meet the conditions from the request
func (h *Handler) GetVenues(ctx context.Context, req *phishqlpb.GetVenuesRequest) (*phishqlpb.GetVenuesResponse, error) {
	log.Printf("%v", req)

	venues, err := h.service.GetVenues(ctx, h.mapper.ProtoToGetVenuesRequest(req))
	if err != nil {
		return nil, err
	}

	return &phishqlpb.GetVenuesResponse{
		Venues: h.mapper.VenuesToProto(venues),
	}, nil
}
