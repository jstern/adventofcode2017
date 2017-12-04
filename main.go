package main

import (
	"fmt"
	"os"

	"github.com/jstern/adventofcode2017/ex1"
)

type exercise func() string

func main() {
	exnum := os.Args[1]

	var ex exercise

	switch exnum {
	case "1":
		ex = ex1.Answer
	default:
		panic("No such exercise")
	}

	answer := ex()

	fmt.Println(answer)
}
