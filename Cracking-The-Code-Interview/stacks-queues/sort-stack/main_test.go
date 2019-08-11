package main

import (
	"testing"
)

func TestSortStack(t *testing.T) {
	var s stack

	s.push(1)
	s.push(2)
	s.push(4)
	s.push(6)
	s.push(8)
	s.push(9)
	s.push(5)
	s.push(3)
	s.push(7)
	s.sort()

	compare := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	if len(s) != len(compare) {
		t.Errorf("\nThis should have the same length")
	}

	for i, v := range compare {
		if s[i] != v {
			t.Errorf("\nThis should value should be: %d, and not %d", v, s[i])
			break
		}
	}
}

func TestMinMax(t *testing.T) {
	var s stack
	var newIndex int
	var expected int
	s.push(4)
	s.push(2)

	newIndex = minMax(&s, 5)
	expected = 2
	if newIndex != expected {
		t.Errorf("\nThis value should be %d, not %d", expected, newIndex)
	}
	newIndex = minMax(&s, 1)
	expected = 0
	if newIndex != expected {
		t.Errorf("\nThis value should be %d, not %d", expected, newIndex)
	}

	newIndex = minMax(&s, 3)
	expected = 1
	if newIndex != expected {
		t.Errorf("\nThis value should be %d, not %d", expected, newIndex)
	}

	var s1 stack

	s1.push(7)
	newIndex = minMax(&s1, 3)
	expected = 0
	if newIndex != expected {
		t.Errorf("\nThis value should be %d, not %d", expected, newIndex)
	}
	newIndex = minMax(&s1, 8)
	expected = 1
	if newIndex != expected {
		t.Errorf("\nThis value should be %d, not %d", expected, newIndex)
	}

}
