package ex8

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReadInputLine(t *testing.T) {
	i := readInputLine("b inc 5 if a > 1")
	assert.Equal(t, "b", i.register)
	assert.Equal(t, "inc", i.op)
	assert.Equal(t, 5, i.amt)
	assert.Equal(t, "a", i.testRegister)
	assert.Equal(t, ">", i.testOp)
	assert.Equal(t, 1, i.testAmt)
}

func TestProcess(t *testing.T) {
	r := map[string]int{}

	i := readInputLine("b inc 5 if a > 1")
	process(i, r)
	assert.Equal(t, 0, r["b"])

	i = readInputLine("a inc 1 if b < 5")
	process(i, r)
	assert.Equal(t, 1, r["a"])

	i = readInputLine("c dec -10 if a >= 1")
	process(i, r)
	assert.Equal(t, 10, r["c"])
}
