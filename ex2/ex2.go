package ex2

import (
	"github.com/jstern/adventofcode2017/tools"
	"math"
	"strconv"
	"strings"
)

type spreadsheet [][]int64

func makeSpreadsheet(input string) spreadsheet {
	lines := strings.Split(input, "\n")
	rows := make([][]int64, len(lines))
	for i, line := range lines {
		cols := strings.Split(line, "\t")
		cells := make([]int64, len(cols))
		for j, col := range cols {
			v, _ := strconv.ParseInt(col, 10, 0)
			cells[j] = v
		}
		rows[i] = cells
	}
	return rows
}

func rowChecksum(row []int64) int64 {
	min := int64(math.MaxInt64)
	max := int64(0)
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

func checksum(sheet spreadsheet) int64 {
	sum := int64(0)
	for _, row := range sheet {
		sum += rowChecksum(row)
	}
	return sum
}

func Answer() string {
	inp, _ := tools.ReadInput("2")
	sheet := makeSpreadsheet(inp)
	cs := checksum(sheet)
	return strconv.FormatInt(cs, 10)
}
