package ex3

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewAllocator(t *testing.T) {
	alloc := newAllocator()
	assert.Equal(t, right, alloc.dir)
	assert.Equal(t, 1, alloc.runLength)
	assert.Equal(t, 0, alloc.runUsed)
	assert.Equal(t, 2, alloc.runsLeftAtLength)
	assert.Equal(t, 0, alloc.dX)
	assert.Equal(t, 0, alloc.dY)
	assert.Equal(t, 0, alloc.distance())
}

func TestAdvance(t *testing.T) {
	alloc := newAllocator()
	alloc.advance()
	assert.Equal(t, 2, alloc.pos)
	assert.Equal(t, up, alloc.dir)
	assert.Equal(t, 1, alloc.dX)
	assert.Equal(t, 0, alloc.dY)
	assert.Equal(t, 1, alloc.distance())
}

func TestAdvanceTo3(t *testing.T) {
	alloc := newAllocator()
	alloc.advanceTo(3)
	assert.Equal(t, 3, alloc.pos)
	assert.Equal(t, left, alloc.dir)
	assert.Equal(t, 1, alloc.dX)
	assert.Equal(t, 1, alloc.dY)
	assert.Equal(t, 2, alloc.distance())
}

func TestAdvanceTo4(t *testing.T) {
	alloc := newAllocator()
	alloc.advanceTo(4)
	assert.Equal(t, left, alloc.dir)
	assert.Equal(t, 0, alloc.dX)
	assert.Equal(t, 1, alloc.dY)
	assert.Equal(t, 1, alloc.distance())
}

func TestAdvanceTo5(t *testing.T) {
	alloc := newAllocator()
	alloc.advanceTo(5)
	assert.Equal(t, down, alloc.dir)
	assert.Equal(t, -1, alloc.dX)
	assert.Equal(t, 1, alloc.dY)
	assert.Equal(t, 2, alloc.distance())
}

func TestAdvanceTo6(t *testing.T) {
	alloc := newAllocator()
	alloc.advanceTo(6)
	assert.Equal(t, down, alloc.dir)
	assert.Equal(t, -1, alloc.dX)
	assert.Equal(t, 0, alloc.dY)
	assert.Equal(t, 1, alloc.distance())
}

func TestAdvanceTo7(t *testing.T) {
	alloc := newAllocator()
	alloc.advanceTo(7)
	assert.Equal(t, right, alloc.dir)
	assert.Equal(t, -1, alloc.dX)
	assert.Equal(t, -1, alloc.dY)
	assert.Equal(t, 2, alloc.distance())
}

func TestAdvanceTo10(t *testing.T) {
	alloc := newAllocator()
	alloc.advanceTo(10)
	assert.Equal(t, up, alloc.dir)
	assert.Equal(t, 2, alloc.dX)
	assert.Equal(t, -1, alloc.dY)
	assert.Equal(t, 3, alloc.distance())
}
