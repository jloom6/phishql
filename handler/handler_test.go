package handler

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	phishqlpb "github.com/jloom6/phishql/.gen/proto/jloom6/phishql"
	"github.com/jloom6/phishql/mapper"
	mmocks "github.com/jloom6/phishql/mapper/mocks"
	smocks "github.com/jloom6/phishql/service/mocks"
	"github.com/jloom6/phishql/structs"
	"github.com/stretchr/testify/assert"
)

func TestHandler_GetShows(t *testing.T) {
	tests := []struct {
		name              string
		getShowsRet       []structs.Show
		getShowsErr       error
		showsToProtoRet   []*phishqlpb.Show
		showsToProtoErr   error
		showsToProtoTimes int
		req               *phishqlpb.GetShowsRequest
		ret               *phishqlpb.GetShowsResponse
		err               error
	}{
		{
			name:        "service.GetShows error",
			getShowsErr: errors.New(""),
			err:         errors.New(""),
		},
		{
			name:              "mapper.ShowsToProto error",
			getShowsRet:       []structs.Show{},
			showsToProtoErr:   errors.New(""),
			showsToProtoTimes: 1,
			err:               errors.New(""),
		},
		{
			name:              "success",
			getShowsRet:       []structs.Show{},
			showsToProtoRet:   []*phishqlpb.Show{},
			showsToProtoTimes: 1,
			ret:               &phishqlpb.GetShowsResponse{Shows: []*phishqlpb.Show{}},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(_t *testing.T) {
			mockCtrl := gomock.NewController(_t)
			defer mockCtrl.Finish()

			mockService := smocks.NewMockInterface(mockCtrl)
			mockMapper := mmocks.NewMockInterface(mockCtrl)

			h := New(Params{
				Service: mockService,
				Mapper:  mockMapper,
			})

			mockMapper.EXPECT().ProtoToGetShowsRequest(&phishqlpb.GetShowsRequest{}).Return(structs.GetShowsRequest{}).Times(1)
			mockService.EXPECT().GetShows(context.Background(), structs.GetShowsRequest{}).Return(test.getShowsRet, test.getShowsErr).Times(1)
			mockMapper.EXPECT().ShowsToProto([]structs.Show{}).Return(test.showsToProtoRet, test.showsToProtoErr).Times(test.showsToProtoTimes)

			ret, err := h.GetShows(context.Background(), &phishqlpb.GetShowsRequest{})

			assert.Equal(_t, test.ret, ret)
			assert.Equal(_t, test.err, err)
		})
	}
}

func TestHandler_GetArtists(t *testing.T) {
	tests := []struct {
		name          string
		getArtistsRet []structs.Artist
		getArtistsErr error
		req           *phishqlpb.GetArtistsRequest
		ret           *phishqlpb.GetArtistsResponse
		err           error
	}{
		{
			name:          "service.GetArtists error",
			getArtistsErr: errors.New(""),
			err:           errors.New(""),
		},
		{
			name:          "success",
			getArtistsRet: []structs.Artist{},
			ret:           &phishqlpb.GetArtistsResponse{Artists: []*phishqlpb.Artist{}},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(_t *testing.T) {
			mockCtrl := gomock.NewController(_t)
			defer mockCtrl.Finish()

			mockService := smocks.NewMockInterface(mockCtrl)

			h := New(Params{
				Service: mockService,
				Mapper:  mapper.New(),
			})

			mockService.EXPECT().GetArtists(context.Background(), structs.GetArtistsRequest{}).Return(test.getArtistsRet, test.getArtistsErr).Times(1)

			ret, err := h.GetArtists(context.Background(), &phishqlpb.GetArtistsRequest{})

			assert.Equal(_t, test.ret, ret)
			assert.Equal(_t, test.err, err)
		})
	}
}

func TestHandler_GetSongs(t *testing.T) {
	tests := []struct {
		name        string
		getSongsRet []structs.Song
		getSongsErr error
		req         *phishqlpb.GetSongsRequest
		ret         *phishqlpb.GetSongsResponse
		err         error
	}{
		{
			name:        "service.GetSongs error",
			getSongsErr: errors.New(""),
			err:         errors.New(""),
		},
		{
			name:        "success",
			getSongsRet: []structs.Song{},
			ret:         &phishqlpb.GetSongsResponse{Songs: []*phishqlpb.Song{}},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(_t *testing.T) {
			mockCtrl := gomock.NewController(_t)
			defer mockCtrl.Finish()

			mockService := smocks.NewMockInterface(mockCtrl)

			h := New(Params{
				Service: mockService,
				Mapper:  mapper.New(),
			})

			mockService.EXPECT().GetSongs(context.Background(), structs.GetSongsRequest{}).Return(test.getSongsRet, test.getSongsErr).Times(1)

			ret, err := h.GetSongs(context.Background(), &phishqlpb.GetSongsRequest{})

			assert.Equal(_t, test.ret, ret)
			assert.Equal(_t, test.err, err)
		})
	}
}

func TestHandler_GetTags(t *testing.T) {
	tests := []struct {
		name       string
		getTagsRet []structs.Tag
		getTagsErr error
		req        *phishqlpb.GetTagsRequest
		ret        *phishqlpb.GetTagsResponse
		err        error
	}{
		{
			name:       "service.GetTags error",
			getTagsErr: errors.New(""),
			err:        errors.New(""),
		},
		{
			name:       "success",
			getTagsRet: []structs.Tag{},
			ret:        &phishqlpb.GetTagsResponse{Tags: []*phishqlpb.Tag{}},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(_t *testing.T) {
			mockCtrl := gomock.NewController(_t)
			defer mockCtrl.Finish()

			mockService := smocks.NewMockInterface(mockCtrl)

			h := New(Params{
				Service: mockService,
				Mapper:  mapper.New(),
			})

			mockService.EXPECT().GetTags(context.Background(), structs.GetTagsRequest{}).Return(test.getTagsRet, test.getTagsErr).Times(1)

			ret, err := h.GetTags(context.Background(), &phishqlpb.GetTagsRequest{})

			assert.Equal(_t, test.ret, ret)
			assert.Equal(_t, test.err, err)
		})
	}
}

func TestHandler_GetTours(t *testing.T) {
	tests := []struct {
		name        string
		getToursRet []structs.Tour
		getToursErr error
		req         *phishqlpb.GetToursRequest
		ret         *phishqlpb.GetToursResponse
		err         error
	}{
		{
			name:        "service.GetTours error",
			getToursErr: errors.New(""),
			err:         errors.New(""),
		},
		{
			name:        "success",
			getToursRet: []structs.Tour{},
			ret:         &phishqlpb.GetToursResponse{Tours: []*phishqlpb.Tour{}},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(_t *testing.T) {
			mockCtrl := gomock.NewController(_t)
			defer mockCtrl.Finish()

			mockService := smocks.NewMockInterface(mockCtrl)

			h := New(Params{
				Service: mockService,
				Mapper:  mapper.New(),
			})

			mockService.EXPECT().GetTours(context.Background(), structs.GetToursRequest{}).Return(test.getToursRet, test.getToursErr).Times(1)

			ret, err := h.GetTours(context.Background(), &phishqlpb.GetToursRequest{})

			assert.Equal(_t, test.ret, ret)
			assert.Equal(_t, test.err, err)
		})
	}
}

func TestHandler_GetVenues(t *testing.T) {
	tests := []struct {
		name         string
		getVenuesRet []structs.Venue
		getVenuesErr error
		req          *phishqlpb.GetVenuesRequest
		ret          *phishqlpb.GetVenuesResponse
		err          error
	}{
		{
			name:         "service.GetVenues error",
			getVenuesErr: errors.New(""),
			err:          errors.New(""),
		},
		{
			name:         "success",
			getVenuesRet: []structs.Venue{},
			ret:          &phishqlpb.GetVenuesResponse{Venues: []*phishqlpb.Venue{}},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(_t *testing.T) {
			mockCtrl := gomock.NewController(_t)
			defer mockCtrl.Finish()

			mockService := smocks.NewMockInterface(mockCtrl)

			h := New(Params{
				Service: mockService,
				Mapper:  mapper.New(),
			})

			mockService.EXPECT().GetVenues(context.Background(), structs.GetVenuesRequest{}).Return(test.getVenuesRet, test.getVenuesErr).Times(1)

			ret, err := h.GetVenues(context.Background(), &phishqlpb.GetVenuesRequest{})

			assert.Equal(_t, test.ret, ret)
			assert.Equal(_t, test.err, err)
		})
	}
}
