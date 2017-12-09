package ex5

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMakeCpu(t *testing.T) {
	inp := "2\n-1"
	c := makeCpu(inp, incr)
	assert.Equal(t, []int{2, -1}, c.instructions)
	assert.Equal(t, 0, c.offset)
}

func TestJump(t *testing.T) {
	c := cpu{[]int{1}, 0, incr}
	c.jump()
	assert.Equal(t, 1, c.offset)
	assert.Equal(t, []int{2}, c.instructions)
}

func TestJumpWithTwiddle(t *testing.T) {
	c := cpu{[]int{1}, 0, twiddle}
	c.jump()
	assert.Equal(t, []int{2}, c.instructions)

	c = cpu{[]int{4}, 0, twiddle}
	c.jump()
	assert.Equal(t, []int{3}, c.instructions)
}

func TestEscapeForward(t *testing.T) {
	c := cpu{[]int{1}, 0, incr}
	c.jump()
	assert.True(t, c.escaped())
}

func TestEscapeBackward(t *testing.T) {
	c := cpu{[]int{-1}, 0, incr}
	c.jump()
	assert.True(t, c.escaped())
}

func TestRun(t *testing.T) {
	c := cpu{[]int{0, 3, 0, 1, -3}, 0, incr}
	steps := c.run()
	assert.Equal(t, 5, steps)
}

func TestRunWithTwiddle(t *testing.T) {
	c := cpu{[]int{0, 3, 0, 1, -3}, 0, twiddle}
	steps := c.run()
	assert.Equal(t, 10, steps)
}
