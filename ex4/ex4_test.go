package ex4

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var x = `
abcde fghij is a valid passphrase.
abcde xyz ecdab is not valid - the letters from the third word can be rearranged to form the first word.
a ab abc abd abf abj is a valid passphrase, because all letters need to be used when forming another word.
iiii oiii ooii oooi oooo is valid.
oiii ioii iioi iiio is not valid - any of these words can be rearranged to form any other word.
`

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
