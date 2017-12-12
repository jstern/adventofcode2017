package ex8

import (
	"fmt"
	"github.com/jstern/adventofcode2017/tools"
	"strconv"
	"strings"
)

type registerMap = map[string]int

type op = func(string, int, registerMap)

var ops = map[string]op{
	"inc": inc,
	"dec": dec,
}

func inc(register string, amt int, registers registerMap) {
	v := registers[register]
	registers[register] = v + amt
}

func dec(register string, amt int, registers registerMap) {
	inc(register, 0-amt, registers)
}

type testFn = func(int, int) bool

var tests = map[string]testFn{
	"==": eq,
	">":  gt,
	">=": gte,
	"!=": neq,
	"<":  lt,
	"<=": lte,
}

func eq(left int, right int) bool {
	return left == right
}

func neq(left int, right int) bool {
	return left != right
}

func gt(left int, right int) bool {
	return left > right
}

func gte(left int, right int) bool {
	return left >= right
}

func lt(left int, right int) bool {
	return left < right
}

func lte(left int, right int) bool {
	return left <= right
}

type instruction struct {
	register     string
	op           string
	amt          int
	testRegister string
	testOp       string
	testAmt      int
}

func readInputLine(line string) instruction {
	parts := strings.Split(line, " ")
	amt, _ := strconv.ParseInt(parts[2], 10, 0)
	testAmt, _ := strconv.ParseInt(parts[6], 10, 0)
	return instruction{
		parts[0],
		parts[1],
		int(amt),
		parts[4],
		parts[5],
		int(testAmt),
	}
}

func instructions(input string) []instruction {
	lines := strings.Split(input, "\n")
	result := make([]instruction, len(lines))
	for i, line := range lines {
		result[i] = readInputLine(line)
	}
	return result
}

func test(instr instruction, registers registerMap) bool {
	return tests[instr.testOp](registers[instr.testRegister], instr.testAmt)
}

func execute(instr instruction, registers registerMap) {
	ops[instr.op](instr.register, instr.amt, registers)
}

func process(instr instruction, registers registerMap) (last int) {
	if test(instr, registers) {
		execute(instr, registers)
		last = registers[instr.register]
	}
	return last
}

func max(registers registerMap) (result int) {
	for _, v := range registers {
		if v > result {
			result = v
		}
	}
	return result
}

func Answer() (result string) {
	result = "?"
	inp, _ := tools.ReadInput("8")
	instructionSet := instructions(inp)
	registers := registerMap{}
	hi := 0
	for _, i := range instructionSet {
		v := process(i, registers)
		if v > hi {
			hi = v
		}
	}
	return fmt.Sprintf("%d, %d", max(registers), hi)
}
