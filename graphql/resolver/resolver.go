package resolver

import (
	"github.com/graphql-go/graphql"
	phishqlpb "github.com/jloom6/phishql/.gen/proto/jloom6/phishql"
	"github.com/jloom6/phishql/mapper"
)

// Interface contains all the resolvers for the graphql schema
type Interface interface {
	GetArtists(p graphql.ResolveParams) (interface{}, error)
	GetShows(p graphql.ResolveParams) (interface{}, error)
	GetSongs(p graphql.ResolveParams) (interface{}, error)
	GetTags(p graphql.ResolveParams) (interface{}, error)
	GetTours(p graphql.ResolveParams) (interface{}, error)
	GetVenues(p graphql.ResolveParams) (interface{}, error)
}

type Resolver struct {
	api    phishqlpb.PhishQLServiceClient
	mapper mapper.Interface
}

type Params struct {
	API    phishqlpb.PhishQLServiceClient
	Mapper mapper.Interface
}

func New(p Params) *Resolver {
	return &Resolver{
		api:    p.API,
		mapper: p.Mapper,
	}
}

func (r *Resolver) GetArtists(p graphql.ResolveParams) (interface{}, error) {
	resp, err := r.api.GetArtists(p.Context, &phishqlpb.GetArtistsRequest{})
	if err != nil {
		return nil, err
	}

	return r.mapper.ProtoToArtists(resp.Artists), nil
}

func (r *Resolver) GetShows(p graphql.ResolveParams) (interface{}, error) {
	req, err := r.mapper.GraphQLShowsToProto(p.Args)
	if err != nil {
		return nil, err
	}

	resp, err := r.api.GetShows(p.Context, req)
	if err != nil {
		return nil, err
	}

	return r.mapper.ProtoToShows(resp.Shows)
}

func (r *Resolver) GetSongs(p graphql.ResolveParams) (interface{}, error) {
	resp, err := r.api.GetSongs(p.Context, &phishqlpb.GetSongsRequest{})
	if err != nil {
		return nil, err
	}

	return r.mapper.ProtoToSongs(resp.Songs), nil
}

func (r *Resolver) GetTags(p graphql.ResolveParams) (interface{}, error) {
	resp, err := r.api.GetTags(p.Context, &phishqlpb.GetTagsRequest{})
	if err != nil {
		return nil, err
	}

	return r.mapper.ProtoToTags(resp.Tags), nil
}

func (r *Resolver) GetTours(p graphql.ResolveParams) (interface{}, error) {
	resp, err := r.api.GetTours(p.Context, &phishqlpb.GetToursRequest{})
	if err != nil {
		return nil, err
	}

	return r.mapper.ProtoToTours(resp.Tours), nil
}

func (r *Resolver) GetVenues(p graphql.ResolveParams) (interface{}, error) {
	resp, err := r.api.GetVenues(p.Context, &phishqlpb.GetVenuesRequest{})
	if err != nil {
		return nil, err
	}

	return r.mapper.ProtoToVenues(resp.Venues), nil
}
