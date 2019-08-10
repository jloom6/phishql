package resolver

//go:generate retool do mockgen -destination=../../.gen/proto/jloom6/phishql/mocks/client.go -package=mocks github.com/jloom6/phishql/.gen/proto/jloom6/phishql PhishQLServiceClient

import (
	"github.com/graphql-go/graphql"
	phishqlpb "github.com/jloom6/phishql/.gen/proto/jloom6/phishql"
	"github.com/jloom6/phishql/mapper"
)

// Interface contains all the resolver functions for the graphql schema
type Interface interface {
	GetArtists(p graphql.ResolveParams) (interface{}, error)
	GetShows(p graphql.ResolveParams) (interface{}, error)
	GetSongs(p graphql.ResolveParams) (interface{}, error)
	GetTags(p graphql.ResolveParams) (interface{}, error)
	GetTours(p graphql.ResolveParams) (interface{}, error)
	GetVenues(p graphql.ResolveParams) (interface{}, error)
}

// Resolver implements the interface for resolving the GraphQL queries
type Resolver struct {
	client phishqlpb.PhishQLServiceClient
	mapper mapper.Interface
}

// Params contains the parameters needed to construct a new resolver
type Params struct {
	Client phishqlpb.PhishQLServiceClient
	Mapper mapper.Interface
}

// New constructs a new resolver
func New(p Params) *Resolver {
	return &Resolver{
		client: p.Client,
		mapper: p.Mapper,
	}
}

// GetArtists gets the artists using the gRPC client
func (r *Resolver) GetArtists(p graphql.ResolveParams) (interface{}, error) {
	resp, err := r.client.GetArtists(p.Context, &phishqlpb.GetArtistsRequest{})
	if err != nil {
		return nil, err
	}

	return r.mapper.ProtoToArtists(resp.Artists), nil
}

// GetShows gets the shows using the gRPC client
func (r *Resolver) GetShows(p graphql.ResolveParams) (interface{}, error) {
	req, err := r.mapper.GraphQLShowsToProto(p.Args)
	if err != nil {
		return nil, err
	}

	resp, err := r.client.GetShows(p.Context, req)
	if err != nil {
		return nil, err
	}

	return r.mapper.ProtoToShows(resp.Shows)
}

// GetSongs gets the songs using the gRPC client
func (r *Resolver) GetSongs(p graphql.ResolveParams) (interface{}, error) {
	resp, err := r.client.GetSongs(p.Context, &phishqlpb.GetSongsRequest{})
	if err != nil {
		return nil, err
	}

	return r.mapper.ProtoToSongs(resp.Songs), nil
}

// GetTags gets the tags using the gRPC client
func (r *Resolver) GetTags(p graphql.ResolveParams) (interface{}, error) {
	resp, err := r.client.GetTags(p.Context, &phishqlpb.GetTagsRequest{})
	if err != nil {
		return nil, err
	}

	return r.mapper.ProtoToTags(resp.Tags), nil
}

// GetTours gets the tours using the gRPC client
func (r *Resolver) GetTours(p graphql.ResolveParams) (interface{}, error) {
	resp, err := r.client.GetTours(p.Context, &phishqlpb.GetToursRequest{})
	if err != nil {
		return nil, err
	}

	return r.mapper.ProtoToTours(resp.Tours), nil
}

// GetVenues gets the venues using the gRPC client
func (r *Resolver) GetVenues(p graphql.ResolveParams) (interface{}, error) {
	resp, err := r.client.GetVenues(p.Context, &phishqlpb.GetVenuesRequest{})
	if err != nil {
		return nil, err
	}

	return r.mapper.ProtoToVenues(resp.Venues), nil
}
