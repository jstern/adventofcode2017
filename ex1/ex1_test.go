package ex1

import "testing"

func TestValueToAdd(t *testing.T) {
	result := valueToAdd(1, 1)
	if result != 1 {
		t.Fail()
	}

	result = valueToAdd(1, 2)
	if result != 0 {
		t.Fail()
	}
}

func TestGetInt(t *testing.T) {
	i := getInt("123", 1)
	if i != 2 {
		t.Fail()
	}
}

func TestValuesToCompare(t *testing.T) {
	f, s := valuesToCompare("32", 0)
	if f != 3 {
		t.Fail()
	}
	if s != 2 {
		t.Fail()
	}
}

func testGetAnswer(t *testing.T) {
	ans := getAnswer("11")
	if ans != 2 {
		t.Fail()
	}

	ans = getAnswer("10")
	if ans != 0 {
		t.Fail()
	}
}
