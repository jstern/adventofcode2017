package ex9

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var readertests = []struct {
	in    string
	score int
}{
	{"{}", 1},
	{"{{{}}}", 6},
	{"{{},{}}", 5},
	{"{{{},{},{{}}}}", 16},
	{"{<a>,<a>,<a>,<a>}", 1},
	{"{{<ab>},{<ab>},{<ab>},{<ab>}}", 9},
	{"{{<!!>},{<!!>},{<!!>},{<!!>}}", 9},
	{"{{<a!>},{<a!>},{<a!>},{<ab>}}", 3},
}

func TestReader(t *testing.T) {
	for _, tt := range readertests {
		r := makeReader(tt.in)
		r.read()
		assert.Equal(t, tt.score, r.total)
	}
}
