package main

import (
	"testing"
)

func TestMinMax(t *testing.T) {
	var s stack
	var newIndex int
	var expected int

	newIndex = minMax(&s, 5)
	expected = -1
	if newIndex != expected {
		t.Errorf("\nThis value should be %d, not %d", expected, newIndex)
	}

	s.push(4)
	expected = 4
	if s[0] != expected {
		t.Errorf("\nThis value should be %d, not %d", expected, s[0])
	}

	s.push(2)
	expected = 2
	if s[0] != expected {
		t.Errorf("\nThis value should be %d, not %d", expected, s[0])
	}

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
	expected = 7
	if s1[0] != expected {
		t.Errorf("\nThis value should be %d, not %d", expected, s1[0])
	}

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

func TestPush(t *testing.T) {
	var s stack
	var expected int

	s.push(7)
	expected = 7
	if s[0] != expected {
		t.Errorf("\nThis value should be %d, not %d", expected, s[0])
	}
}
