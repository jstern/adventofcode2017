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

var garbagetests = []struct {
	in    string
	total int
}{
	{"<>", 0},
	{"<random characters>", 17},
	{"<<<<>", 3},
	{"<{!>}>", 2},
	{"<!!>", 0},
	{"<!!!>>", 0},
	{"<{o\"i!a,<{i<a>", 10},
}

func TestGarbage(t *testing.T) {
	for _, tt := range garbagetests {
		r := makeReader(tt.in)
		r.read()
		assert.Equal(t, tt.total, r.garbageTotal)
	}
}
