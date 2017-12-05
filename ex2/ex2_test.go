package ex2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var inp = `5	1	9	5
7	5	3
2	4	6	8`

var sheet = makeSpreadsheet(inp)

func TestMakeSpreadsheet(t *testing.T) {
	assert.Equal(t, len(sheet), 3, "made a sheet with 3 rows")
	assert.Equal(
		t,
		[]int64{5, 1, 9, 5},
		sheet[0],
		"made a sheet with expected cells",
	)
}

func TestRowChecksum(t *testing.T) {
	assert.Equal(t, int64(8), rowChecksum(sheet[0]))
	assert.Equal(t, int64(4), rowChecksum(sheet[1]))
	assert.Equal(t, int64(6), rowChecksum(sheet[2]))
}

func TestChecksum(t *testing.T) {
	assert.Equal(t, int64(18), checksum(sheet))
}
