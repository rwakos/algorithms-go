package main

import (
	"testing"
)

func TestIsCircular(t *testing.T) {
	values := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	l1 := new(LinkedListNode)
	node := new(LinkedListNode)
	for x := range values {
		l1.addNode(x)
	}

	if isCircular(l1) {
		t.Errorf("\nFAIL, this should be false")
	}

	i := 0
	for l := l1; l != nil; l = l.next {
		if i == 5 {
			node = l
		}

		if i == 8 {
			l.next = node
			break
		}
		i++
	}

	if !isCircular(l1) {
		t.Errorf("\nFAIL, this should be true")
	}
}
