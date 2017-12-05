package ex1

import (
	"github.com/jstern/adventofcode2017/tools"
	"strconv"
)

type puzzle struct {
	input string
	gap   int
}

func (puz puzzle) index(pos int) int {
	max := len(puz.input)
	if pos >= max {
		return pos - max
	}
	return pos
}

func (puz puzzle) valueAt(pos int) int64 {
	c := string([]byte{puz.input[puz.index(pos)]})
	res, err := strconv.ParseInt(c, 10, 0)
	if err != nil {
		panic("on the streets of london")
	}
	return res
}

func (puz puzzle) calculateAnswer() int64 {
	a := int64(0)
	for i := 0; i < len(puz.input); i++ {
		cur := puz.valueAt(i)
		next := puz.valueAt(i + puz.gap)
		if cur == next {
			a += cur
		}
	}
	return a
}

func Answer() string {
	inp, _ := tools.ReadInput("1")
	puz := puzzle{inp, 1}
	a := puz.calculateAnswer()
	return strconv.FormatInt(a, 10)
}

func Answer2() string {
	inp, _ := tools.ReadInput("1")
	gap := len(inp) / 2
	puz := puzzle{inp, gap}
	a := puz.calculateAnswer()
	return strconv.FormatInt(a, 10)
}
