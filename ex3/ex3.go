package ex3

import (
	"strconv"
)

var example = `
36  35  34  33  32  31
17  16  15  14  13  30
18   5   4   3  12  29
19   6   1   2  11  28
20   7   8   9  10  27
21  22  23  24  25  26
`

// 1 1 2 2 3 3

var right = 0
var up = 1
var left = 2
var down = 3

var dirs = []int{right, up, left, down}

type allocator struct {
	pos              int
	dir              int
	runLength        int
	runUsed          int
	runsLeftAtLength int
	dX               int
	dY               int
}

func newAllocator() *allocator {
	return &allocator{1, right, 1, 0, 2, 0, 0}
}

func (a *allocator) turn() {
	next := a.dir + 1
	if next > down {
		next = 0
	}
	a.dir = dirs[next]
}

func (a *allocator) advance() {
	a.pos++
	a.runUsed++
	if a.dir == right {
		a.dX++
	}
	if a.dir == left {
		a.dX--
	}
	if a.dir == up {
		a.dY++
	}
	if a.dir == down {
		a.dY--
	}
	if a.runUsed == a.runLength {
		a.turn()
		a.runsLeftAtLength--
		if a.runsLeftAtLength == 0 {
			a.runsLeftAtLength = 2
			a.runLength++
		}
		a.runUsed = 0
	}
}

func (a *allocator) advanceTo(n int) {
	for i := a.pos; a.pos < n; i++ {
		a.advance()
	}
}

func (a *allocator) distance() int {
	return abs(a.dX) + abs(a.dY)
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func Answer() string {
	a := newAllocator()
	a.advanceTo(277678)
	return strconv.FormatInt(int64(a.distance()), 10)
}
