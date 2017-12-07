package ex4

import (
	"github.com/jstern/adventofcode2017/tools"
	"sort"
	"strconv"
	"strings"
)

type validationWord func(string) string

func passphrases(input string) []string {
	return strings.Split(input, "\n")
}

func noop(word string) string {
	return word
}

func letters(word string) string {
	lets := strings.Split(word, "")
	sort.Strings(lets)
	return strings.Join(lets, "")
}

func valid(phrase string, fn validationWord) bool {
	seen := make(map[string]string)
	words := strings.Split(phrase, " ")
	for _, word := range words {
		vw := fn(word)
		if seen[vw] != "" {
			return false
		}
		seen[vw] = word
	}
	return true
}

func answer(fn validationWord) int {
	inp, _ := tools.ReadInput("4")
	phrases := passphrases(inp)
	total := 0
	for _, phrase := range phrases {
		if valid(phrase, fn) {
			total++
		}
	}
	return total
}

func Answer() string {
	total := answer(noop)
	return strconv.FormatInt(int64(total), 10)
}

func Answer2() string {
	total := answer(letters)
	return strconv.FormatInt(int64(total), 10)
}
