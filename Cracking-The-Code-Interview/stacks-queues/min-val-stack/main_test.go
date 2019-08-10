package main

import (
	"testing"
)

func TestPop(t *testing.T) {
	var f FullStack

	f.push(10)
	f.push(21)
	f.push(7)
	if f.min != 7 {
		t.Errorf("\nThis should be equal to 7, got: %d", f.min)
	}
	f.push(4)
	f.push(3)
	if f.min != 3 {
		t.Errorf("\nThis should be equal to 3, got: %d", f.min)
	}
	f.pop()
	if f.min != 4 {
		t.Errorf("\nThis should be equal to 4, got: %d", f.min)
	}
}
