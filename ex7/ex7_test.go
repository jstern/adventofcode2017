package ex7

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReadInputLineNoChildren(t *testing.T) {
	name, weight, children := readInputLine("twix (2)")
	assert.Equal(t, "twix", name)
	assert.Equal(t, 2, weight)
	assert.Equal(t, []string{}, children)
}

func TestReadInputLineWithChild(t *testing.T) {
	name, weight, children := readInputLine("twix (2) -> zagnut")
	assert.Equal(t, "twix", name)
	assert.Equal(t, 2, weight)
	assert.Equal(t, []string{"zagnut"}, children)
}

func TestReadInputLineWithChildren(t *testing.T) {
	name, weight, children := readInputLine("twix (20) -> zagnut, zero")
	assert.Equal(t, "twix", name)
	assert.Equal(t, 20, weight)
	assert.Equal(t, []string{"zagnut", "zero"}, children)
}

func TestWeighDiscs(t *testing.T) {
	prog := &program{"twix", 1, 0, []*program{}}
	weighDiscs(prog)
	assert.Equal(t, 0, prog.discWeight)

	zag := &program{"zagnut", 2, 0, []*program{}}
	twix := &program{"twix", 1, 0, []*program{zag}}
	weighDiscs(twix)
	assert.Equal(t, 2, twix.discWeight)

	zag = &program{"zagnut", 2, 0, []*program{}}
	zero := &program{"zero", 2, 0, []*program{}}
	twix = &program{"twix", 1, 0, []*program{zag, zero}}
	weighDiscs(twix)
	assert.Equal(t, 4, twix.discWeight)
	assert.Equal(t, 0, zag.discWeight)
	assert.Equal(t, 0, zero.discWeight)

	zero = &program{"zero", 2, 0, []*program{}}
	zag = &program{"zagnut", 2, 0, []*program{zero}}
	twix = &program{"twix", 1, 0, []*program{zag}}
	weighDiscs(twix)
	assert.Equal(t, 4, twix.discWeight)
	assert.Equal(t, 2, zag.discWeight)
	assert.Equal(t, 0, zero.discWeight)
}

func TestDiscBalanced(t *testing.T) {
	zag := &program{"zagnut", 2, 0, []*program{}}
	zero := &program{"zero", 1, 0, []*program{}}
	twix := &program{"twix", 1, 0, []*program{zag, zero}}
	weighDiscs(twix)
	assert.True(t, discBalanced(*zag))
	assert.True(t, discBalanced(*zero))
	assert.False(t, discBalanced(*twix))
}
