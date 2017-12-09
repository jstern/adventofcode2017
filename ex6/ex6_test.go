package ex6

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFullestBank(t *testing.T) {
	banks := []int{0, 5, 10, 0, 10}
	assert.Equal(t, 2, fullestBank(banks))
}

func TestNextBank(t *testing.T) {
	assert.Equal(t, 1, nextBank(0, 2))
	assert.Equal(t, 0, nextBank(1, 2))
}

func TestReallocate(t *testing.T) {
	banks := []int{0, 1}

	reallocate(banks)
	assert.Equal(t, []int{1, 0}, banks)

	reallocate(banks)
	assert.Equal(t, []int{0, 1}, banks)

	banks = []int{0, 2}
	reallocate(banks)
	assert.Equal(t, []int{1, 1}, banks)
}

func TestAllocation(t *testing.T) {
	banks := []int{0, 1}
	assert.Equal(t, "[0 1]", allocation(banks))
}

func TestCycles(t *testing.T) {
	banks := []int{0, 1}
	total, loop := cycles(banks)
	assert.Equal(t, 2, total)
	assert.Equal(t, 2, loop)

	banks = []int{0, 2, 7, 0}
	total, loop = cycles(banks)
	assert.Equal(t, 5, total)
	assert.Equal(t, 4, loop)
}
