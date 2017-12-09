package ex6

import (
	"fmt"
	"strconv"
)

func fullestBank(banks []int) int {
	high := 0
	indices := []int{0}
	for i, s := range banks {
		if s > high {
			high = s
			indices = []int{i}
		}
		if s == high {
			indices = append(indices, i)
		}
	}
	return indices[0]
}

func nextBank(cur int, ct int) int {
	i := cur + 1
	if i < ct {
		return i
	}
	return ct - i
}

func reallocate(banks []int) {
	ct := len(banks)
	from := fullestBank(banks)
	blocks := banks[from]
	banks[from] = 0
	next := nextBank(from, ct)
	for i := 0; i < blocks; i++ {
		banks[next]++
		next = nextBank(next, ct)
	}
}

func allocation(banks []int) string {
	return fmt.Sprint(banks)
}

func cycles(banks []int) (int, int) {
	i := 0
	s := 0
	allocations := make(map[string]int)
	allocations[allocation(banks)] = -1
	for {
		reallocate(banks)
		i++
		allocation := allocation(banks)
		s = allocations[allocation]
		if s != 0 {
			if s == -1 {
				s = 0
			}
			break
		}
		allocations[allocation] = i
	}
	return i, i - s
}

func Answer() string {
	banks := []int{0, 5, 10, 0, 11, 14, 13, 4, 11, 8, 8, 7, 1, 4, 12, 11}
	total, _ := cycles(banks)
	return strconv.FormatInt(int64(total), 10)
}

func Answer2() string {
	banks := []int{0, 5, 10, 0, 11, 14, 13, 4, 11, 8, 8, 7, 1, 4, 12, 11}
	_, loop := cycles(banks)
	return strconv.FormatInt(int64(loop), 10)
}
