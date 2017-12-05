package ex2

import (
	"github.com/jstern/adventofcode2017/tools"
	"math"
	"sort"
	"strconv"
	"strings"
)

type spreadsheet [][]int

type rowChecksum func([]int) int

func makeSpreadsheet(input string) spreadsheet {
	lines := strings.Split(input, "\n")
	rows := make([][]int, len(lines))
	for i, line := range lines {
		cols := strings.Split(line, "\t")
		cells := make([]int, len(cols))
		for j, col := range cols {
			v, _ := strconv.ParseInt(col, 10, 0)
			cells[j] = int(v)
		}
		rows[i] = cells
	}
	return rows
}

func rowChecksum1(row []int) int {
	min := math.MaxInt64
	max := 0
	for _, c := range row {
		if c > max {
			max = c
		}
		if c < min {
			min = c
		}
	}
	return max - min
}

func rowChecksum2(row []int) int {
	// sort first to make it likely to hit earlier
	sorted := make([]int, len(row))
	copy(sorted, row)
	sort.Ints(sorted)
	for i, v := range sorted {
		for c := i + 1; c < len(sorted); c++ {
			v2 := sorted[c]
			if v2%v == 0 {
				return v2 / v
			}
		}
	}
	panic("bad row!")
}

func checksum(sheet spreadsheet, f rowChecksum) int {
	sum := 0
	for _, row := range sheet {
		sum += f(row)
	}
	return sum
}

func Answer() string {
	inp, _ := tools.ReadInput("2")
	sheet := makeSpreadsheet(inp)
	cs := checksum(sheet, rowChecksum1)
	return strconv.FormatInt(int64(cs), 10)
}

func Answer2() string {
	inp, _ := tools.ReadInput("2")
	sheet := makeSpreadsheet(inp)
	cs := checksum(sheet, rowChecksum2)
	return strconv.FormatInt(int64(cs), 10)
}
