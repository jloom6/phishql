package handler

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	phishqlpb "github.com/jloom6/phishql/.gen/proto/jloom6/phishql"
	mmocks "github.com/jloom6/phishql/mapper/mocks"
	smocks "github.com/jloom6/phishql/service/mocks"
	"github.com/jloom6/phishql/structs"
	"github.com/stretchr/testify/assert"
)

func TestHandler_GetShows(t *testing.T) {
	tests := []struct{
		name string
		getShowsRet []structs.Show
		getShowsErr error
		showsToProtoRet []*phishqlpb.Show
		showsToProtoErr error
		showsToProtoTimes int
		req *phishqlpb.GetShowsRequest
		ret *phishqlpb.GetShowsResponse
		err error
	}{
		{
			name: "service.GetShows error",
			getShowsErr: errors.New(""),
			err: errors.New(""),
		},
		{
			name: "mapper.ShowsToProto error",
			getShowsRet: []structs.Show{},
			showsToProtoErr: errors.New(""),
			showsToProtoTimes: 1,
			err: errors.New(""),
		},
		{
			name: "success",
			getShowsRet: []structs.Show{},
			showsToProtoRet: []*phishqlpb.Show{},
			showsToProtoTimes: 1,
			ret: &phishqlpb.GetShowsResponse{Shows:[]*phishqlpb.Show{}},
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
				Mapper: mockMapper,
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
