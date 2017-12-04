package ex1

import "testing"

func TestIndex(t *testing.T) {
	puz := puzzle{"11", 1}

	i := puz.index(0)
	if i != 0 {
		t.Fail()
	}

	i = puz.index(1)
	if i != 1 {
		t.Fail()
	}

	i = puz.index(2)
	if i != 0 {
		t.Fail()
	}
}

func TestValueAt(t *testing.T) {
	puz := puzzle{"12", 1}

	i := puz.valueAt(0)
	if i != 1 {
		t.Fail()
	}

	i = puz.valueAt(1)
	if i != 2 {
		t.Fail()
	}

	i = puz.valueAt(2)
	if i != 1 {
		t.Fail()
	}
}

func testCalculateAnswer(t *testing.T) {
	puz := puzzle{"1122", 1}
	ans := puz.calculateAnswer()
	if ans != 3 {
		t.Fail()
	}

	puz = puzzle{"1111", 1}
	ans = puz.calculateAnswer()
	if ans != 4 {
		t.Fail()
	}

	puz = puzzle{"1212", 2}
	ans = puz.calculateAnswer()
	if ans != 6 {
		t.Fail()
	}
}
