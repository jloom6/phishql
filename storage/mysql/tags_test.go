package mysql

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/jloom6/phishql/internal/db/mocks"
	"github.com/jloom6/phishql/structs"
	"github.com/stretchr/testify/assert"
)

func TestStore_GetTags(t *testing.T) {
	dest := []interface{}{gomock.Any(), gomock.Any()}

	tag := structs.Tag{
		ID:   6,
		Text: "Simpsons Signal",
	}

	tests := []struct {
		name               string
		queryErr           error
		rowsCloseTimes     int
		rowsNextTrueTimes  int
		rowsNextFalseTimes int
		rowsScanDo         func(...interface{})
		rowsScanTimes      int
		ret                []structs.Tag
		err                error
	}{
		{
			name:     "db.QueryContext error",
			queryErr: errors.New(""),
			err:      errors.New(""),
		},
		{
			name:               "success",
			rowsCloseTimes:     1,
			rowsNextTrueTimes:  1,
			rowsNextFalseTimes: 1,
			rowsScanDo: func(dest ...interface{}) {
				*(dest[0].(*int)) = tag.ID
				*(dest[1].(*string)) = tag.Text
			},
			rowsScanTimes: 1,
			ret:           []structs.Tag{tag},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(_t *testing.T) {
			mockCtrl := gomock.NewController(_t)
			defer mockCtrl.Finish()

			mockDB := mocks.NewMockInterface(mockCtrl)
			mockRows := mocks.NewMockRows(mockCtrl)

			mockDB.EXPECT().QueryContext(context.Background(), gomock.Any()).Return(mockRows, test.queryErr).Times(1)
			mockRows.EXPECT().Close().Return(nil).Times(test.rowsCloseTimes)
			gomock.InOrder(
				mockRows.EXPECT().Next().Return(true).Times(test.rowsNextTrueTimes),
				mockRows.EXPECT().Next().Return(false).Times(test.rowsNextFalseTimes),
			)
			mockRows.EXPECT().Scan(dest...).Do(test.rowsScanDo).Return(nil).Times(test.rowsScanTimes)

			s := New(Params{
				DB: mockDB,
			})

			ret, err := s.GetTags(context.Background(), structs.GetTagsRequest{})

			assert.Equal(_t, test.ret, ret)
			assert.Equal(_t, test.err, err)
		})
	}
}

func TestMakeTags(t *testing.T) {
	dest := []interface{}{gomock.Any(), gomock.Any()}

	tag := structs.Tag{
		ID:   2,
		Text: "Vacuum Solo",
	}

	tests := []struct {
		name           string
		nextTrueTimes  int
		nextFalseTimes int
		scanDo         func(...interface{})
		scanErr        error
		ret            []structs.Tag
		err            error
	}{
		{
			name:          "rows.Scan error",
			nextTrueTimes: 1,
			scanDo:        func(...interface{}) {},
			scanErr:       errors.New(""),
			err:           errors.New(""),
		},
		{
			name:           "success",
			nextTrueTimes:  1,
			nextFalseTimes: 1,
			scanDo: func(dest ...interface{}) {
				*(dest[0].(*int)) = tag.ID
				*(dest[1].(*string)) = tag.Text
			},
			ret: []structs.Tag{tag},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(_t *testing.T) {
			mockCtrl := gomock.NewController(_t)
			defer mockCtrl.Finish()

			mockRows := mocks.NewMockRows(mockCtrl)

			gomock.InOrder(
				mockRows.EXPECT().Next().Return(true).Times(test.nextTrueTimes),
				mockRows.EXPECT().Next().Return(false).Times(test.nextFalseTimes),
			)

			mockRows.EXPECT().Scan(dest...).Do(test.scanDo).Return(test.scanErr).Times(1)

			ret, err := makeTags(mockRows)

			assert.Equal(_t, test.ret, ret)
			assert.Equal(_t, test.err, err)
		})
	}
}
