package mysql

import (
	"testing"

	"code.uber.internal/infra/uown/go-build/.go/src/gb2/src/github.com/magiconair/properties/assert"
	"github.com/jloom6/phishql/structs"
)

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
