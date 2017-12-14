package ex9

import (
	"fmt"
	"github.com/jstern/adventofcode2017/tools"
	"strings"
)

type reader struct {
	stream       []string
	max          int
	pos          int
	level        int
	total        int
	garbageCur   int
	garbageTotal int
	ignore       bool
	inGarbage    bool
	done         bool
}

func makeReader(stream string) *reader {
	chars := strings.Split(stream, "")
	return &reader{chars, len(chars), 0, 0, 0, 0, 0, false, false, false}
}

func (r *reader) next() {
	if r.pos >= r.max {
		r.done = true
		return
	}

	c := r.stream[r.pos]
	r.pos++

	if r.ignore {
		r.ignore = false
		return
	}

	switch c {
	case "!":
		r.ignore = true
	case "<":
		if r.inGarbage {
			r.garbageCur++
		}
		r.inGarbage = true
	case ">":
		r.inGarbage = false
		r.garbageTotal += r.garbageCur
		r.garbageCur = 0
	case "{":
		if !r.inGarbage {
			r.level++
		}
	case "}":
		if !r.inGarbage {
			r.total += r.level
			r.level--
		}
	default:
		// nothing
	}
	if r.inGarbage && !r.ignore && c != "<" {
		r.garbageCur++
	}
}

func (r *reader) read() {
	for {
		if r.done {
			return
		}
		r.next()
	}
}

func Answer() string {
	inp, _ := tools.ReadInput("9")
	r := makeReader(inp)
	r.read()
	return fmt.Sprintf("%d, %d", r.total, r.garbageTotal)
}
