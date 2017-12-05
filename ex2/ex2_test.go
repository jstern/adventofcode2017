package ex2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var inp = `5	1	9	5
7	5	3
2	4	6	8`

var inp2 = `5	9	2	8
9	4	7	3
3	8	6	5`

var sheet = makeSpreadsheet(inp)
var sheet2 = makeSpreadsheet(inp2)

func TestMakeSpreadsheet(t *testing.T) {
	assert.Equal(t, len(sheet), 3, "made a sheet with 3 rows")
	assert.Equal(
		t,
		[]int{5, 1, 9, 5},
		sheet[0],
		"made a sheet with expected cells",
	)
}

func TestRowChecksum1(t *testing.T) {
	assert.Equal(t, 8, rowChecksum1(sheet[0]))
	assert.Equal(t, 4, rowChecksum1(sheet[1]))
	assert.Equal(t, 6, rowChecksum1(sheet[2]))
}

func TestRowChecksum2(t *testing.T) {
	assert.Equal(t, 4, rowChecksum2(sheet2[0]))
	assert.Equal(t, 3, rowChecksum2(sheet2[1]))
	assert.Equal(t, 2, rowChecksum2(sheet2[2]))
}

func TestChecksum(t *testing.T) {
	assert.Equal(t, 18, checksum(sheet, rowChecksum1))
	assert.Equal(t, 9, checksum(sheet2, rowChecksum2))
}
