package ex4

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPhraseWithDuplicateWordsIsInvalid(t *testing.T) {
	assert.False(t, valid("aa bb aa", noop))
	assert.False(t, valid("aa bb cc dd aa", noop))
}

func TestPhraseWithoutDuplicateWordsIsInvalid(t *testing.T) {
	assert.True(t, valid("aa bb cc", noop))
	assert.True(t, valid("aa bb cc dd ee", noop))
	assert.True(t, valid("aa bb cc dd ee", noop))
	assert.True(t, valid("aa bb cc dd aaa", noop))
}

func TestPassphrases(t *testing.T) {
	assert.Equal(t, []string{"aa", "bb"}, passphrases("aa\nbb"))
}

func TestPhraseWithAnagramsIsInvalid(t *testing.T) {
	assert.False(t, valid("abcde xyz ecdab", letters))
}
