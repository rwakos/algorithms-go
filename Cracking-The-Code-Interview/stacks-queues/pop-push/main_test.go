package main

import "testing"

func TestPop(t *testing.T) {
	var s stack
	_, err := s.pop()
	if !err {
		t.Errorf("\nThis should throw an error")
	}

	s.push(5)

	if s[0] != 5 {
		t.Errorf("\nThis should be a stack of one item equal to 5")
	}

	s.push(7)

	if s[0] != 7 || s[1] != 5 {
		t.Errorf("\nThis should be a stack of one item equal to 7 > 5")
	}
}
