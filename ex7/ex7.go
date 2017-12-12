package ex7

import (
	"fmt"
	"github.com/jstern/adventofcode2017/tools"
	"regexp"
	"strconv"
	"strings"
)

var programRe = regexp.MustCompile("(.*?) \\((.*?)\\)")

type program struct {
	name       string
	weight     int
	discWeight int
	disc       []*program
}

func readInputLine(line string) (name string, weight int, children []string) {
	parts := strings.Split(line, " -> ")
	prgParts := programRe.FindAllStringSubmatch(parts[0], -1)
	name = prgParts[0][1]
	wgt, _ := strconv.ParseInt(prgParts[0][2], 10, 0)
	weight = int(wgt)
	if len(parts) == 2 {
		children = strings.Split(parts[1], ", ")
	} else {
		children = []string{}
	}
	return name, weight, children
}

func readData(inp string) (root *program) {
	lines := strings.Split(inp, "\n")
	programs := make(map[string]*program)
	childrenOf := make(map[string][]string)
	isChild := make(map[string]bool)
	for _, line := range lines {
		name, weight, children := readInputLine(line)
		programs[name] = &program{
			name,
			weight,
			0,
			make([]*program, len(children)),
		}
		childrenOf[name] = children
		for _, child := range children {
			isChild[child] = true
		}
	}
	for name, prog := range programs {
		if !isChild[name] {
			root = prog
		}
		for i, child := range childrenOf[name] {
			prog.disc[i] = programs[child]
		}
	}
	return root
}

func weighDiscs(prog *program) int {
	for _, child := range prog.disc {
		prog.discWeight += weighDiscs(child)
		prog.discWeight += child.weight
	}
	return prog.discWeight
}

func same(disc []*program, fn func(p program) int) (result bool) {
	value := -1
	for _, p := range disc {
		v := fn(*p)
		if value != -1 && v != value {
			return false
		}
		value = v
	}
	return true
}

func discWeight(p program) int {
	return p.discWeight
}

func towerWeight(p program) int {
	return p.weight + p.discWeight
}

func discBalanced(p program) bool {
	return same(p.disc, towerWeight)
}

func trace(p *program) (res []program) {
	fmt.Printf("%+v\n", p)
	res = []program{*p}
	for _, child := range p.disc {
		balanced := discBalanced(*child)
		if !balanced {
			res = append(res, trace(child)...)
		}
	}
	return res
}

func findCulprit(unbalanced *program) (result *program) {
	badPath := trace(unbalanced)
	end := badPath[len(badPath)-1]
	fmt.Printf("%s %d %d\n", end.name, end.weight, towerWeight(end))
	for _, c := range end.disc {
		fmt.Printf("  %s %d %d\n", c.name, c.weight, towerWeight(*c))
	}
	return &end
}

func Answer() (result string) {
	result = "?"
	inp, _ := tools.ReadInput("7")
	root := readData(inp)
	return root.name
}

func Answer2() (result string) {
	result = "?"
	inp, _ := tools.ReadInput("7")
	root := readData(inp)
	weighDiscs(root)
	findCulprit(root)
	return "yeah"
}

func Example() (result string) {
	result = "?"
	inp, _ := tools.ReadInput("7e")
	root := readData(inp)
	weighDiscs(root)
	findCulprit(root)
	return "yeah"
}
