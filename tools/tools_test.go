package tools

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReadsInputFromFile(t *testing.T) {
	res, err := ReadInput("0")
	assert.Equal(t, res, "hi everybody", "It reads a string.")
	assert.Nil(t, err)
}

func TestReturnsErrForBadFile(t *testing.T) {
	res, err := ReadInput("blah")
	assert.Equal(t, res, "", "It returns an empty string.")
	assert.NotNil(t, err)
}
