package service

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/jloom6/phishql/storage/mocks"
	"github.com/jloom6/phishql/structs"
	"github.com/stretchr/testify/assert"
)

func TestService_GetShows(t *testing.T) {
	show1 := structs.Show{Date:time.Now()}
	show2 := structs.Show{Date:show1.Date.Add(-1 * time.Hour)}

	tests := []struct{
		name string
		getShowsRet []structs.Show
		getShowsErr error
		ret []structs.Show
		err error
	}{
		{
			name: "store.GetShows error",
			getShowsErr: errors.New(""),
			err: errors.New(""),
		},
		{
			name: "success",
			getShowsRet: []structs.Show{show1, show2},
			ret: []structs.Show{show2, show1},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(_t *testing.T) {
			mockCtrl := gomock.NewController(_t)
			defer mockCtrl.Finish()

			mockStore := mocks.NewMockInterface(mockCtrl)

			s := New(Params{
				Store:mockStore,
			})

			mockStore.EXPECT().GetShows(context.Background(), structs.GetShowsRequest{}).Return(test.getShowsRet, test.getShowsErr).Times(1)

			ret, err := s.GetShows(context.Background(), structs.GetShowsRequest{})

			assert.Equal(_t, test.ret, ret)
			assert.Equal(_t, test.err, err)
		})
	}
}

func TestService_GetArtists(t *testing.T) {
	tests := []struct{
		name string
		getArtistsRet []structs.Artist
		getArtistErr error
		req structs.GetArtistsRequest
		ret []structs.Artist
		err error
	}{
		{
			name: "success",
			getArtistsRet: []structs.Artist{},
			req: structs.GetArtistsRequest{},
			ret: []structs.Artist{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(_t *testing.T) {
			mockCtrl := gomock.NewController(_t)
			defer mockCtrl.Finish()

			mockStore := mocks.NewMockInterface(mockCtrl)

			s := New(Params{
				Store:mockStore,
			})

			mockStore.EXPECT().GetArtists(context.Background(), structs.GetArtistsRequest{}).Return(test.getArtistsRet, test.getArtistErr).Times(1)

			ret, err := s.GetArtists(context.Background(), structs.GetArtistsRequest{})

			assert.Equal(_t, test.ret, ret)
			assert.Equal(_t, test.err, err)
		})
	}
}

