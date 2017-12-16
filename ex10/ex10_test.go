package ex10

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRingIndex(t *testing.T) {
	h := makeHasher(3)
	assert.Equal(t, 0, h.ringIndex(0))
	assert.Equal(t, 1, h.ringIndex(1))
	assert.Equal(t, 2, h.ringIndex(2))
	assert.Equal(t, 0, h.ringIndex(3))
	assert.Equal(t, 1, h.ringIndex(4))
	assert.Equal(t, 2, h.ringIndex(5))
	assert.Equal(t, 0, h.ringIndex(6))
}

func TestTie(t *testing.T) {
	h := makeHasher(5)
	h.tie(3)
	assert.Equal(t, []int{2, 1, 0, 3, 4}, h.digits)
	assert.Equal(t, 3, h.pos)
	assert.Equal(t, 1, h.skip)

	h.tie(4)
	assert.Equal(t, []int{4, 3, 0, 1, 2}, h.digits)
	assert.Equal(t, 3, h.pos)
	assert.Equal(t, 2, h.skip)
}

func TestConvertInput(t *testing.T) {
	actual := convertInput("1,2,3")
	expected := []int{49, 44, 50, 44, 51, 17, 31, 73, 47, 23}
	assert.Equal(t, expected, actual)
}

func TestCondense(t *testing.T) {
	sparse := []int{65, 27, 9, 1, 4, 3, 40, 50, 91, 7, 6, 0, 2, 5, 68, 22}
	condensed := condenseHash(sparse)
	assert.Equal(t, []int{64}, condensed)
}

func TestHexlify(t *testing.T) {
	hash := []int{9, 99, 249}
	assert.Equal(t, "0963f9", hexlify(hash))
}
