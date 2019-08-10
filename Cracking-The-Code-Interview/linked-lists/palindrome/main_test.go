package main

import (
	"fmt"
	"testing"
)

func TestReverseList(t *testing.T) {
	values := []int{7, 1, 6, 1, 7}
	l1 := &LinkedListNode{}
	for _, v := range values {
		l1.addNode(v)
	}
	l2 := l1

	compare := l1.isEqualTo(l2)

	if !compare {
		fmt.Println("FAIL")
		t.Errorf("\nThis should be a palindrome")
	}
}
