package ex5

import (
	"github.com/jstern/adventofcode2017/tools"
	"strconv"
	"strings"
)

type frob func(int) int

type cpu struct {
	instructions []int
	offset       int
	update       frob
}

func incr(v int) int {
	return v + 1
}

func twiddle(v int) int {
	if v >= 3 {
		return v - 1
	}
	return v + 1
}

func (c cpu) next() int {
	return c.offset + c.instructions[c.offset]
}

func (c cpu) escaped() bool {
	if c.offset < 0 || c.offset >= len(c.instructions) {
		return true
	}
	return false
}

func (c *cpu) jump() {
	dest := c.next()
	c.instructions[c.offset] = c.update(c.instructions[c.offset])
	c.offset = dest
}

func makeCpu(inp string, f frob) *cpu {
	lines := strings.Split(inp, "\n")
	instructions := make([]int, len(lines))
	for i, s := range lines {
		v, _ := strconv.ParseInt(s, 10, 0)
		instructions[i] = int(v)
	}
	return &cpu{instructions, 0, f}
}

func (c *cpu) run() int {
	jumps := 0
	for {
		c.jump()
		jumps++
		if c.escaped() {
			break
		}
	}
	return jumps
}

func Answer() string {
	inp, _ := tools.ReadInput("5")
	cpu := makeCpu(inp, incr)
	result := cpu.run()
	return strconv.FormatInt(int64(result), 10)
}

func Answer2() string {
	inp, _ := tools.ReadInput("5")
	cpu := makeCpu(inp, twiddle)
	result := cpu.run()
	return strconv.FormatInt(int64(result), 10)
}
