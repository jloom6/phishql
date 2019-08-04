package mysql

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/jloom6/phishql/internal/db/mocks"
	"github.com/jloom6/phishql/structs"
	"github.com/stretchr/testify/assert"
)

func TestStore_GetShows(t *testing.T) {
	showDest := []interface{}{gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(),
		gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(),gomock.Any(), gomock.Any(),
		gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()}
	setDest := []interface{}{gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()}
	songDest := []interface{}{gomock.Any(), gomock.Any(), gomock.Any(),
		gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()}

	show := structs.Show{
		ID: 420,
		Date: time.Now(),
		Artist: structs.Artist{
			ID: 2,
			Name: "Pork Tornado",
		},
		Venue: structs.Venue{
			ID: 3,
			Name: "The Gorge",
			City: "George",
			State: "WA",
			Country: "USA",
		},
		Tour: &structs.Tour{
			ID: 4,
			Name: "Summer Tour 2018",
			Description: "Set Your Soul Free becomes the go to opener",
		},
		Notes: "This show was good.",
		Soundcheck: "Foam",
		Sets: []structs.Set{
			{
				ID: 1,
				Label: "SET 1",
				Songs: []structs.SetSong{
					{
						Song: structs.Song{
							ID: 555,
							Name: "Mock Song",
						},
						Tag: &structs.Tag{
							ID: 1,
							Text: "Debut",
						},
						Transition: "->",
					},
				},
			},
		},
	}

	tests := []struct{
		name string
		getShowsErr error
		showRowsCloseTimes int
		showRowsNextTrueTimes int
		showRowsNextFalseTimes int
		showRowsScanDo func(...interface{})
		showRowsScanTimes int
		getSetsErr error
		getSetsTimes int
		setRowsCloseTimes int
		setRowsNextTrueTimes int
		setRowsNextFalseTimes int
		setRowsScanDo func(...interface{})
		setRowsScanTimes int
		getSongsErr error
		getSongsTimes int
		songRowsCloseTimes int
		songRowsNextTrueTimes int
		songRowsNextFalseTimes int
		songRowsScanDo func(...interface{})
		songRowsScanTimes int
		ret []structs.Show
		err error
	}{
		{
			name: "s.getShows error",
			getShowsErr: errors.New(""),
			err: errors.New(""),
		},
		{
			name: "s.getSets error",
			showRowsCloseTimes: 1,
			showRowsNextFalseTimes: 1,
			getSetsErr: errors.New(""),
			getSetsTimes: 1,
			err: errors.New(""),
		},
		{
			name: "s.getSongs error",
			showRowsCloseTimes: 1,
			showRowsNextFalseTimes: 1,
			getSetsTimes: 1,
			setRowsCloseTimes: 1,
			setRowsNextFalseTimes: 1,
			getSongsErr: errors.New(""),
			getSongsTimes: 1,
			err: errors.New(""),
		},
		{
			name: "success",
			showRowsCloseTimes: 1,
			showRowsNextTrueTimes: 1,
			showRowsNextFalseTimes: 1,
			showRowsScanDo: func(dest ...interface{}) {
				*(dest[0].(*int)) = show.ID
				*(dest[1].(*time.Time)) = show.Date
				*(dest[2].(*int)) = show.Artist.ID
				*(dest[3].(*string)) = show.Artist.Name
				*(dest[4].(*int)) = show.Venue.ID
				*(dest[5].(*string)) = show.Venue.Name
				*(dest[6].(*string)) = show.Venue.City
				*(dest[7].(*string)) = show.Venue.State
				*(dest[8].(*string)) = show.Venue.Country
				*(dest[9].(*int)) = show.Tour.ID
				*(dest[10].(*string)) = show.Tour.Name
				*(dest[11].(*string)) = show.Tour.Description
				*(dest[12].(*string)) = show.Notes
				*(dest[13].(*string)) = show.Soundcheck
			},
			showRowsScanTimes: 1,
			getSetsTimes: 1,
			setRowsCloseTimes: 1,
			setRowsNextTrueTimes: 1,
			setRowsNextFalseTimes: 1,
			setRowsScanDo: func(dest ...interface{}) {
				*(dest[0].(*int)) = show.ID
				*(dest[1].(*int)) = 0
				*(dest[2].(*int)) = show.Sets[0].ID
				*(dest[3].(*string)) = show.Sets[0].Label
			},
			setRowsScanTimes: 1,
			getSongsTimes: 1,
			songRowsCloseTimes: 1,
			songRowsNextTrueTimes: 1,
			songRowsNextFalseTimes: 1,
			songRowsScanDo: func(dest ...interface{}) {
				*(dest[0].(*int)) = show.ID
				*(dest[1].(*int)) = 0
				*(dest[2].(*int)) = show.Sets[0].Songs[0].Song.ID
				*(dest[3].(*string)) = show.Sets[0].Songs[0].Song.Name
				*(dest[4].(*int)) = show.Sets[0].Songs[0].Tag.ID
				*(dest[5].(*string)) = show.Sets[0].Songs[0].Tag.Text
				*(dest[6].(*string)) = show.Sets[0].Songs[0].Transition
			},
			songRowsScanTimes: 1,
			ret: []structs.Show{show},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(_t *testing.T) {
			mockCtrl := gomock.NewController(_t)
			defer mockCtrl.Finish()

			mockDB := mocks.NewMockInterface(mockCtrl)
			mockShowRows := mocks.NewMockRows(mockCtrl)
			mockSetRows := mocks.NewMockRows(mockCtrl)
			mockSongRows := mocks.NewMockRows(mockCtrl)

			// getShows
			mockDB.EXPECT().QueryContext(context.Background(), gomock.Any(), gomock.Any()).Return(mockShowRows, test.getShowsErr).Times(1)
			mockShowRows.EXPECT().Close().Return(nil).Times(test.showRowsCloseTimes)
			gomock.InOrder(
				mockShowRows.EXPECT().Next().Return(true).Times(test.showRowsNextTrueTimes),
				mockShowRows.EXPECT().Next().Return(false).Times(test.showRowsNextFalseTimes),
			)
			mockShowRows.EXPECT().Scan(showDest...).Do(test.showRowsScanDo).Return(nil).Times(test.showRowsScanTimes)

			// getSets
			mockDB.EXPECT().QueryContext(context.Background(), gomock.Any(), gomock.Any()).Return(mockSetRows, test.getSetsErr).Times(test.getSetsTimes)
			mockSetRows.EXPECT().Close().Return(nil).Times(test.setRowsCloseTimes)
			gomock.InOrder(
				mockSetRows.EXPECT().Next().Return(true).Times(test.setRowsNextTrueTimes),
				mockSetRows.EXPECT().Next().Return(false).Times(test.setRowsNextFalseTimes),
			)
			mockSetRows.EXPECT().Scan(setDest...).Do(test.setRowsScanDo).Return(nil).Times(test.setRowsScanTimes)

			// getSongs
			mockDB.EXPECT().QueryContext(context.Background(), gomock.Any(), gomock.Any()).Return(mockSongRows, test.getSongsErr).Times(test.getSongsTimes)
			mockSongRows.EXPECT().Close().Return(nil).Times(test.songRowsCloseTimes)
			gomock.InOrder(
				mockSongRows.EXPECT().Next().Return(true).Times(test.songRowsNextTrueTimes),
				mockSongRows.EXPECT().Next().Return(false).Times(test.songRowsNextFalseTimes),
			)
			mockSongRows.EXPECT().Scan(songDest...).Do(test.songRowsScanDo).Return(nil).Times(test.songRowsScanTimes)

			s := New(Params{
				DB: mockDB,
			})

			ret, err := s.GetShows(context.Background(), structs.GetShowsRequest{})

			assert.Equal(_t, test.ret, ret)
			assert.Equal(_t, test.err, err)
		})
	}
}

func TestMakeClauseAndArgs(t *testing.T) {
	tests := []struct{
		name string
		condition structs.Condition
		clause string
		args []interface{}
	}{
		{
			name: "and or base",
			condition: structs.Condition{
				Ands: []structs.Condition{
					{
						Ors: []structs.Condition{
							{
								Base: structs.BaseCondition{
									DayOfWeek: 1,
								},
							},
							{
								Base: structs.BaseCondition{
									DayOfWeek: 7,
								},
							},
						},
					},
					{
						Base: structs.BaseCondition{
							State: "NJ",
						},
					},
				},
			},
			clause: `((
CASE WHEN ? = 0 THEN 1=1 ELSE YEAR(shows.date) = ? END AND
CASE WHEN ? = 0 THEN 1=1 ELSE MONTH(shows.date) = ? END AND
CASE WHEN ? = 0 THEN 1=1 ELSE DAY(shows.date) = ? END AND
CASE WHEN ? = 0 THEN 1=1 ELSE DAYOFWEEK(shows.date) = ? END AND
CASE WHEN ? = '' THEN 1=1 ELSE venues.city = ? END AND
CASE WHEN ? = '' THEN 1=1 ELSE venues.state = ? END AND
CASE WHEN ? = '' THEN 1=1 ELSE venues.country = ? END
) OR (
CASE WHEN ? = 0 THEN 1=1 ELSE YEAR(shows.date) = ? END AND
CASE WHEN ? = 0 THEN 1=1 ELSE MONTH(shows.date) = ? END AND
CASE WHEN ? = 0 THEN 1=1 ELSE DAY(shows.date) = ? END AND
CASE WHEN ? = 0 THEN 1=1 ELSE DAYOFWEEK(shows.date) = ? END AND
CASE WHEN ? = '' THEN 1=1 ELSE venues.city = ? END AND
CASE WHEN ? = '' THEN 1=1 ELSE venues.state = ? END AND
CASE WHEN ? = '' THEN 1=1 ELSE venues.country = ? END
)) AND (
CASE WHEN ? = 0 THEN 1=1 ELSE YEAR(shows.date) = ? END AND
CASE WHEN ? = 0 THEN 1=1 ELSE MONTH(shows.date) = ? END AND
CASE WHEN ? = 0 THEN 1=1 ELSE DAY(shows.date) = ? END AND
CASE WHEN ? = 0 THEN 1=1 ELSE DAYOFWEEK(shows.date) = ? END AND
CASE WHEN ? = '' THEN 1=1 ELSE venues.city = ? END AND
CASE WHEN ? = '' THEN 1=1 ELSE venues.state = ? END AND
CASE WHEN ? = '' THEN 1=1 ELSE venues.country = ? END
)`,
			args: []interface{}{
				0, 0, 0, 0, 0, 0, 1, 1, "", "", "", "", "", "",
				0, 0, 0, 0, 0, 0, 7, 7, "", "", "", "", "", "",
				0, 0, 0, 0, 0, 0, 0, 0, "", "", "NJ", "NJ", "", "",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(_t *testing.T) {
			clause, args := makeClauseAndArgs(test.condition)

			assert.Equal(_t, test.clause, clause)
			assert.Equal(_t, test.args, args)
		})
	}
}

func TestMakeShows(t *testing.T) {
	dest := []interface{}{gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(),
		gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(),gomock.Any(), gomock.Any(),
		gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()}

	show := structs.Show{
		ID: 420,
		Date: time.Now(),
		Artist: structs.Artist{
			ID: 2,
			Name: "Pork Tornado",
		},
		Venue: structs.Venue{
			ID: 3,
			Name: "The Gorge",
			City: "George",
			State: "WA",
			Country: "USA",
		},
		Tour: &structs.Tour{
			ID: 4,
			Name: "Summer Tour 2018",
			Description: "Set Your Soul Free becomes the go to opener",
		},
		Notes: "This show was good.",
		Soundcheck: "Foam",
	}

	tests := []struct{
		name string
		nextTrueTimes int
		nextFalseTimes int
		scanDo func(...interface{})
		scanErr error
		ret map[int]structs.Show
		err error
	}{
		{
			name: "rows.Scan error",
			nextTrueTimes: 1,
			scanDo: func(...interface{}) {},
			scanErr: errors.New(""),
			err: errors.New(""),
		},
		{
			name: "success",
			nextTrueTimes: 1,
			nextFalseTimes: 1,
			scanDo: func(dest ...interface{}) {
				*(dest[0].(*int)) = show.ID
				*(dest[1].(*time.Time)) = show.Date
				*(dest[2].(*int)) = show.Artist.ID
				*(dest[3].(*string)) = show.Artist.Name
				*(dest[4].(*int)) = show.Venue.ID
				*(dest[5].(*string)) = show.Venue.Name
				*(dest[6].(*string)) = show.Venue.City
				*(dest[7].(*string)) = show.Venue.State
				*(dest[8].(*string)) = show.Venue.Country
				*(dest[9].(*int)) = show.Tour.ID
				*(dest[10].(*string)) = show.Tour.Name
				*(dest[11].(*string)) = show.Tour.Description
				*(dest[12].(*string)) = show.Notes
				*(dest[13].(*string)) = show.Soundcheck
			},
			ret: map[int]structs.Show{
				show.ID: show,
			},
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

			ret, err := makeShows(mockRows)

			assert.Equal(_t, test.ret, ret)
			assert.Equal(_t, test.err, err)
		})
	}
}

func TestMakeSets(t *testing.T) {
	dest := []interface{}{gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()}

	showID := 420
	setOrder := 0

	set := structs.Set{
		ID: 1,
		Label: "SET 1",
	}

	tests := []struct{
		name string
		nextTrueTimes int
		nextFalseTimes int
		scanDo func(...interface{})
		scanErr error
		ret map[int]map[int]structs.Set
		err error
	}{
		{
			name: "rows.Scan error",
			nextTrueTimes: 1,
			scanDo: func(...interface{}) {},
			scanErr: errors.New(""),
			err: errors.New(""),
		},
		{
			name: "success",
			nextTrueTimes: 1,
			nextFalseTimes: 1,
			scanDo: func(dest ...interface{}) {
				*(dest[0].(*int)) = showID
				*(dest[1].(*int)) = setOrder
				*(dest[2].(*int)) = set.ID
				*(dest[3].(*string)) = set.Label
			},
			ret: map[int]map[int]structs.Set{
				showID: {
					setOrder: set,
				},
			},
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

			ret, err := makeSets(mockRows)

			assert.Equal(_t, test.ret, ret)
			assert.Equal(_t, test.err, err)
		})
	}
}

func TestMakeSongs(t *testing.T) {
	dest := []interface{}{gomock.Any(), gomock.Any(), gomock.Any(),
		gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()}

	showID := 420
	setOrder := 0

	setSong := structs.SetSong{
		Song: structs.Song{
			ID: 555,
			Name: "Mock Song",
		},
		Tag: &structs.Tag{
			ID: 1,
			Text: "Debut",
		},
		Transition: "->",
	}

	tests := []struct{
		name string
		nextTrueTimes int
		nextFalseTimes int
		scanDo func(...interface{})
		scanErr error
		ret map[int]map[int][]structs.SetSong
		err error
	}{
		{
			name: "rows.Scan error",
			nextTrueTimes: 1,
			scanDo: func(...interface{}) {},
			scanErr: errors.New(""),
			err: errors.New(""),
		},
		{
			name: "success",
			nextTrueTimes: 1,
			nextFalseTimes: 1,
			scanDo: func(dest ...interface{}) {
				*(dest[0].(*int)) = showID
				*(dest[1].(*int)) = setOrder
				*(dest[2].(*int)) = setSong.Song.ID
				*(dest[3].(*string)) = setSong.Song.Name
				*(dest[4].(*int)) = setSong.Tag.ID
				*(dest[5].(*string)) = setSong.Tag.Text
				*(dest[6].(*string)) = setSong.Transition
			},
			ret: map[int]map[int][]structs.SetSong{
				showID: {
					setOrder: {setSong},
				},
			},
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

			ret, err := makeSongs(mockRows)

			assert.Equal(_t, test.ret, ret)
			assert.Equal(_t, test.err, err)
		})
	}
}

func TestHydrateSet(t *testing.T) {
	tests := []struct{
		name string
		sets map[int]structs.Set
		songs map[int][]structs.SetSong
		idx int
		set structs.Set
		hasMore bool
	}{
		{
			name: "sets[idx] doesn't exist",
			sets: map[int]structs.Set{},
		},
		{
			name: "songs is nil",
			sets: map[int]structs.Set{0: {}},
			hasMore: true,
		},
		{
			name: "songs[idx] doesn't exist",
			sets: map[int]structs.Set{0: {}},
			songs: map[int][]structs.SetSong{},
			hasMore: true,
		},
		{
			name: "success",
			sets: map[int]structs.Set{0: {}},
			songs: map[int][]structs.SetSong{0: {}},
			set: structs.Set{Songs:[]structs.SetSong{}},
			hasMore: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(_t *testing.T) {
			set, hasMore := hydrateSet(test.sets, test.songs, test.idx)

			assert.Equal(_t, test.set, set)
			assert.Equal(_t, test.hasMore, hasMore)
		})
	}
}
