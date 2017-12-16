package ex10

import (
	"fmt"
	"strings"
)

var input = []int{
	88, 88, 211, 106, 141, 1, 78, 254, 2, 111, 77, 255, 90, 0, 54, 205,
}

var input2 = "88,88,211,106,141,1,78,254,2,111,77,255,90,0,54,205"

var tail = []int{17, 31, 73, 47, 23}

type knotHasher struct {
	digits []int
	pos    int
	skip   int
}

func makeHasher(size int) *knotHasher {
	digits := make([]int, size)
	for i := 0; i < size; i++ {
		digits[i] = i
	}

	return &knotHasher{
		digits,
		0,
		0,
	}
}

func (h *knotHasher) ringIndex(i int) int {
	return i % len(h.digits)
}

func (h *knotHasher) reverseSegment(length int) {
	end := h.pos + length - 1
	reversed := make([]int, length)
	for i := 0; i < length; i++ {
		reversed[i] = h.digits[h.ringIndex(end-i)]
	}
	for i, v := range reversed {
		h.digits[h.ringIndex(h.pos+i)] = v
	}
}

func (h *knotHasher) tie(length int) {
	h.reverseSegment(length)
	h.pos = h.ringIndex(h.pos + h.skip + length)
	h.skip++
}

func convertInput(input string) []int {
	result := make([]int, len(input)+5)
	bytes := []byte(input)
	for i, v := range bytes {
		result[i] = int(v)
	}
	for i, v := range tail {
		result[len(result)-5+i] = v
	}
	return result
}

func xor(hash []int, start int) int {
	result := hash[start]
	for i := 1; i < 16; i++ {
		result = result ^ hash[i+start]
	}
	return result
}

func condenseHash(hash []int) []int {
	result := make([]int, len(hash)/16)
	buckets := len(hash) / 16
	for i := 0; i < buckets; i++ {
		result[i] = xor(hash, i*16)
	}
	return result
}

func hexlify(hash []int) string {
	hex := make([]string, len(hash))
	for i, v := range hash {
		hex[i] = fmt.Sprintf("%02x", v)
	}
	return strings.Join(hex, "")
}

func Answer() string {
	h := makeHasher(256)
	for _, v := range input {
		h.tie(v)
	}
	return fmt.Sprintf("%d", h.digits[0]*h.digits[1])
}

func Answer2() string {
	h := makeHasher(256)
	input := convertInput(input2)
	for i := 0; i < 64; i++ {
		for _, v := range input {
			h.tie(v)
		}
	}
	condensed := condenseHash(h.digits)
	return fmt.Sprintf("%v", hexlify(condensed))
}
