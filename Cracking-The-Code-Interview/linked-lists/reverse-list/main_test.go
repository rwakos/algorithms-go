package main

import (
	"fmt"
	"testing"
)

func TestReverseList(t *testing.T) {
	values := []int{7, 1, 6}
	l1 := &LinkedListNode{}
	for _, v := range values {
		l1.addNode(v)
	}

	compare := "->6->1->7"
	l1.reverseList()
	result := l1.toString()
	if result != compare {
		fmt.Println("FAIL")
		t.Errorf("\nThis should be: %s, and not: %s", compare, result)
	}
}
