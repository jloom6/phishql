package resolver

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/graphql-go/graphql"
	phishqlpb "github.com/jloom6/phishql/.gen/proto/jloom6/phishql"
	cmocks "github.com/jloom6/phishql/.gen/proto/jloom6/phishql/mocks"
	mmocks "github.com/jloom6/phishql/mapper/mocks"
	"github.com/jloom6/phishql/structs"
	"github.com/stretchr/testify/assert"
)

func TestResolver_GetArtists(t *testing.T) {
	tests := []struct {
		name                string
		getArtistsRet       *phishqlpb.GetArtistsResponse
		getArtistsErr       error
		protoToArtistsRet   []structs.Artist
		protoToArtistsTimes int
		params              graphql.ResolveParams
		ret                 interface{}
		err                 error
	}{
		{
			name: "client.GetArtists error",
			params: graphql.ResolveParams{
				Context: context.Background(),
			},
			getArtistsErr: errors.New("some error"),
			err:           errors.New("some error"),
		},
		{
			name: "success",
			params: graphql.ResolveParams{
				Context: context.Background(),
			},
			getArtistsRet:       &phishqlpb.GetArtistsResponse{},
			protoToArtistsRet:   []structs.Artist{},
			protoToArtistsTimes: 1,
			ret:                 []structs.Artist{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(_t *testing.T) {
			mockCtrl := gomock.NewController(_t)
			defer mockCtrl.Finish()

			mockClient := cmocks.NewMockPhishQLServiceClient(mockCtrl)
			mockMapper := mmocks.NewMockInterface(mockCtrl)

			r := New(Params{
				Client: mockClient,
				Mapper: mockMapper,
			})

			mockClient.EXPECT().GetArtists(context.Background(), &phishqlpb.GetArtistsRequest{}).Return(test.getArtistsRet, test.getArtistsErr).Times(1)
			mockMapper.EXPECT().ProtoToArtists(gomock.Any()).Return(test.protoToArtistsRet).Times(test.protoToArtistsTimes)

			ret, err := r.GetArtists(test.params)

			assert.Equal(_t, test.ret, ret)
			assert.Equal(_t, test.err, err)
		})
	}
}

func TestResolver_GetShows(t *testing.T) {
	tests := []struct {
		name               string
		showsReqToProtoRet *phishqlpb.GetShowsRequest
		showsReqToProtoErr error
		getShowsRet        *phishqlpb.GetShowsResponse
		getShowsErr        error
		getShowsTimes      int
		protoToShowsRet    []structs.Show
		protoToShowsErr    error
		protoToShowsTimes  int
		params             graphql.ResolveParams
		ret                interface{}
		err                error
	}{
		{
			name: "mapper.GraphQLShowsToProto error",
			params: graphql.ResolveParams{
				Context: context.Background(),
			},
			showsReqToProtoErr: errors.New("some error"),
			err:                errors.New("some error"),
		},
		{
			name: "client.GetShows error",
			params: graphql.ResolveParams{
				Context: context.Background(),
			},
			showsReqToProtoRet: &phishqlpb.GetShowsRequest{},
			getShowsErr:        errors.New("some error"),
			getShowsTimes:      1,
			err:                errors.New("some error"),
		},
		{
			name: "success",
			params: graphql.ResolveParams{
				Context: context.Background(),
			},
			showsReqToProtoRet: &phishqlpb.GetShowsRequest{},
			getShowsRet:        &phishqlpb.GetShowsResponse{},
			getShowsTimes:      1,
			protoToShowsRet:    []structs.Show{},
			protoToShowsTimes:  1,
			ret:                []structs.Show{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(_t *testing.T) {
			mockCtrl := gomock.NewController(_t)
			defer mockCtrl.Finish()

			mockClient := cmocks.NewMockPhishQLServiceClient(mockCtrl)
			mockMapper := mmocks.NewMockInterface(mockCtrl)

			r := New(Params{
				Client: mockClient,
				Mapper: mockMapper,
			})

			mockMapper.EXPECT().GraphQLShowsToProto(gomock.Any()).Return(test.showsReqToProtoRet, test.showsReqToProtoErr).Times(1)
			mockClient.EXPECT().GetShows(context.Background(), &phishqlpb.GetShowsRequest{}).Return(test.getShowsRet, test.getShowsErr).Times(test.getShowsTimes)
			mockMapper.EXPECT().ProtoToShows(gomock.Any()).Return(test.protoToShowsRet, test.protoToShowsErr).Times(test.protoToShowsTimes)

			ret, err := r.GetShows(test.params)

			assert.Equal(_t, test.ret, ret)
			assert.Equal(_t, test.err, err)
		})
	}
}

func TestResolver_GetSongs(t *testing.T) {
	tests := []struct {
		name             string
		getSongsRet      *phishqlpb.GetSongsResponse
		getSongsErr      error
		protoToSongsRet  []structs.Song
		protoToSongTimes int
		params           graphql.ResolveParams
		ret              interface{}
		err              error
	}{
		{
			name: "client.GetSongs error",
			params: graphql.ResolveParams{
				Context: context.Background(),
			},
			getSongsErr: errors.New("some error"),
			err:         errors.New("some error"),
		},
		{
			name: "success",
			params: graphql.ResolveParams{
				Context: context.Background(),
			},
			getSongsRet:      &phishqlpb.GetSongsResponse{},
			protoToSongsRet:  []structs.Song{},
			protoToSongTimes: 1,
			ret:              []structs.Song{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(_t *testing.T) {
			mockCtrl := gomock.NewController(_t)
			defer mockCtrl.Finish()

			mockClient := cmocks.NewMockPhishQLServiceClient(mockCtrl)
			mockMapper := mmocks.NewMockInterface(mockCtrl)

			r := New(Params{
				Client: mockClient,
				Mapper: mockMapper,
			})

			mockClient.EXPECT().GetSongs(context.Background(), &phishqlpb.GetSongsRequest{}).Return(test.getSongsRet, test.getSongsErr).Times(1)
			mockMapper.EXPECT().ProtoToSongs(gomock.Any()).Return(test.protoToSongsRet).Times(test.protoToSongTimes)

			ret, err := r.GetSongs(test.params)

			assert.Equal(_t, test.ret, ret)
			assert.Equal(_t, test.err, err)
		})
	}
}

func TestResolver_GetTags(t *testing.T) {
	tests := []struct {
		name             string
		getTagsRet       *phishqlpb.GetTagsResponse
		getTagsErr       error
		protoToTagsRet   []structs.Tag
		protoToTagsTimes int
		params           graphql.ResolveParams
		ret              interface{}
		err              error
	}{
		{
			name: "client.GetTags error",
			params: graphql.ResolveParams{
				Context: context.Background(),
			},
			getTagsErr: errors.New("some error"),
			err:        errors.New("some error"),
		},
		{
			name: "success",
			params: graphql.ResolveParams{
				Context: context.Background(),
			},
			getTagsRet:       &phishqlpb.GetTagsResponse{},
			protoToTagsRet:   []structs.Tag{},
			protoToTagsTimes: 1,
			ret:              []structs.Tag{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(_t *testing.T) {
			mockCtrl := gomock.NewController(_t)
			defer mockCtrl.Finish()

			mockClient := cmocks.NewMockPhishQLServiceClient(mockCtrl)
			mockMapper := mmocks.NewMockInterface(mockCtrl)

			r := New(Params{
				Client: mockClient,
				Mapper: mockMapper,
			})

			mockClient.EXPECT().GetTags(context.Background(), &phishqlpb.GetTagsRequest{}).Return(test.getTagsRet, test.getTagsErr).Times(1)
			mockMapper.EXPECT().ProtoToTags(gomock.Any()).Return(test.protoToTagsRet).Times(test.protoToTagsTimes)

			ret, err := r.GetTags(test.params)

			assert.Equal(_t, test.ret, ret)
			assert.Equal(_t, test.err, err)
		})
	}
}

func TestResolver_GetTours(t *testing.T) {
	tests := []struct {
		name              string
		getToursRet       *phishqlpb.GetToursResponse
		getToursErr       error
		protoToToursRet   []structs.Tour
		protoToToursTimes int
		params            graphql.ResolveParams
		ret               interface{}
		err               error
	}{
		{
			name: "client.GetTours error",
			params: graphql.ResolveParams{
				Context: context.Background(),
			},
			getToursErr: errors.New("some error"),
			err:         errors.New("some error"),
		},
		{
			name: "success",
			params: graphql.ResolveParams{
				Context: context.Background(),
			},
			getToursRet:       &phishqlpb.GetToursResponse{},
			protoToToursRet:   []structs.Tour{},
			protoToToursTimes: 1,
			ret:               []structs.Tour{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(_t *testing.T) {
			mockCtrl := gomock.NewController(_t)
			defer mockCtrl.Finish()

			mockClient := cmocks.NewMockPhishQLServiceClient(mockCtrl)
			mockMapper := mmocks.NewMockInterface(mockCtrl)

			r := New(Params{
				Client: mockClient,
				Mapper: mockMapper,
			})

			mockClient.EXPECT().GetTours(context.Background(), &phishqlpb.GetToursRequest{}).Return(test.getToursRet, test.getToursErr).Times(1)
			mockMapper.EXPECT().ProtoToTours(gomock.Any()).Return(test.protoToToursRet).Times(test.protoToToursTimes)

			ret, err := r.GetTours(test.params)

			assert.Equal(_t, test.ret, ret)
			assert.Equal(_t, test.err, err)
		})
	}
}

func TestResolver_GetVenues(t *testing.T) {
	tests := []struct {
		name               string
		getVenuesRet       *phishqlpb.GetVenuesResponse
		getVenuesErr       error
		protoToVenuesRet   []structs.Venue
		protoToVenuesTimes int
		params             graphql.ResolveParams
		ret                interface{}
		err                error
	}{
		{
			name: "client.GetVenues error",
			params: graphql.ResolveParams{
				Context: context.Background(),
			},
			getVenuesErr: errors.New("some error"),
			err:          errors.New("some error"),
		},
		{
			name: "success",
			params: graphql.ResolveParams{
				Context: context.Background(),
			},
			getVenuesRet:       &phishqlpb.GetVenuesResponse{},
			protoToVenuesRet:   []structs.Venue{},
			protoToVenuesTimes: 1,
			ret:                []structs.Venue{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(_t *testing.T) {
			mockCtrl := gomock.NewController(_t)
			defer mockCtrl.Finish()

			mockClient := cmocks.NewMockPhishQLServiceClient(mockCtrl)
			mockMapper := mmocks.NewMockInterface(mockCtrl)

			r := New(Params{
				Client: mockClient,
				Mapper: mockMapper,
			})

			mockClient.EXPECT().GetVenues(context.Background(), &phishqlpb.GetVenuesRequest{}).Return(test.getVenuesRet, test.getVenuesErr).Times(1)
			mockMapper.EXPECT().ProtoToVenues(gomock.Any()).Return(test.protoToVenuesRet).Times(test.protoToVenuesTimes)

			ret, err := r.GetVenues(test.params)

			assert.Equal(_t, test.ret, ret)
			assert.Equal(_t, test.err, err)
		})
	}
}
