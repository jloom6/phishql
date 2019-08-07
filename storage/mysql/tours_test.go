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

func TestStore_GetTours(t *testing.T) {
	dest := []interface{}{gomock.Any(), gomock.Any(), gomock.Any()}

	tour := structs.Tour{
		ID:          1997,
		Name:        "1997 Fall Tour",
		Description: "Phish Destroys America",
	}

	tests := []struct {
		name               string
		queryErr           error
		rowsCloseTimes     int
		rowsNextTrueTimes  int
		rowsNextFalseTimes int
		rowsScanDo         func(...interface{})
		rowsScanTimes      int
		ret                []structs.Tour
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
				*(dest[0].(*int)) = tour.ID
				*(dest[1].(*string)) = tour.Name
				*(dest[2].(*string)) = tour.Description
			},
			rowsScanTimes: 1,
			ret:           []structs.Tour{tour},
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

			ret, err := s.GetTours(context.Background(), structs.GetToursRequest{})

			assert.Equal(_t, test.ret, ret)
			assert.Equal(_t, test.err, err)
		})
	}
}

func TestMakeTours(t *testing.T) {
	dest := []interface{}{gomock.Any(), gomock.Any(), gomock.Any()}

	tour := structs.Tour{
		ID:          1998,
		Name:        "The Island Tour",
		Description: "Phish got bored",
	}

	tests := []struct {
		name           string
		nextTrueTimes  int
		nextFalseTimes int
		scanDo         func(...interface{})
		scanErr        error
		ret            []structs.Tour
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
				*(dest[0].(*int)) = tour.ID
				*(dest[1].(*string)) = tour.Name
				*(dest[2].(*string)) = tour.Description
			},
			ret: []structs.Tour{tour},
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

			ret, err := makeTours(mockRows)

			assert.Equal(_t, test.ret, ret)
			assert.Equal(_t, test.err, err)
		})
	}
}
