package main

import (
	"fmt"
	"os"

	"github.com/jstern/adventofcode2017/ex1"
	"github.com/jstern/adventofcode2017/ex2"
	"github.com/jstern/adventofcode2017/ex3"
	"github.com/jstern/adventofcode2017/ex4"
)

type exercise func() string

func main() {
	exnum := os.Args[1]

	var ex exercise

	switch exnum {
	case "1":
		ex = ex1.Answer
	case "1a":
		ex = ex1.Answer2
	case "2":
		ex = ex2.Answer
	case "2a":
		ex = ex2.Answer2
	case "3":
		ex = ex3.Answer
	case "3a":
		ex = ex3.Answer2
	case "4":
		ex = ex4.Answer
	case "4a":
		ex = ex4.Answer2
	default:
		panic("No such exercise")
	}

	answer := ex()

	fmt.Println(answer)
}
