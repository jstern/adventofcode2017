package ex3

import (
	"strconv"
)

var right = 0
var up = 1
var left = 2
var down = 3

var dirs = [4]int{right, up, left, down}
var turns = [4]int{up, left, down, right}
var dxs = [4]int{1, 0, -1, 0}
var dys = [4]int{0, 1, 0, -1}

type coord struct {
	x int
	y int
}

type allocator struct {
	pos              int
	dir              int
	runLength        int
	runUsed          int
	runsLeftAtLength int
	x                int
	y                int
	values           map[coord]int
}

func newAllocator() *allocator {
	values := make(map[coord]int)
	values[coord{0, 0}] = 1
	return &allocator{1, right, 1, 0, 2, 0, 0, values}
}

func (a *allocator) turn() {
	a.dir = turns[a.dir]
}

func (a *allocator) advance() {
	a.runUsed++
	a.x += dxs[a.dir]
	a.y += dys[a.dir]
	if a.runUsed == a.runLength {
		a.turn()
		a.runsLeftAtLength--
		if a.runsLeftAtLength == 0 {
			a.runsLeftAtLength = 2
			a.runLength++
		}
		a.runUsed = 0
	}
	a.pos++
	a.setValue()
}

func (a *allocator) advanceTo(n int) {
	for i := a.pos; a.pos < n; i++ {
		a.advance()
	}
}

func (a allocator) distance() int {
	return abs(a.x) + abs(a.y)
}

func (a *allocator) setValue() {
	v := 0
	for x := -1; x < 2; x++ {
		for y := -1; y < 2; y++ {
			// this will include the current position,
			// but that's because its value is 0
			// (assuming this is called once)
			v += a.values[coord{a.x + x, a.y + y}]
		}
	}
	a.values[coord{a.x, a.y}] = v
}

func (a allocator) lastValue() int {
	return a.values[coord{a.x, a.y}]
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

func Answer2() string {
	a := newAllocator()
	v := 0
	for {
		a.advance()
		v = a.lastValue()
		if v > 277678 {
			break
		}
	}
	return strconv.FormatInt(int64(v), 10)
}
