package main

import (
	"fmt"
	"testing"
)

func TestSumLists(t *testing.T) {
	values := []int{7, 1, 6}
	l1 := &LinkedListNode{}
	l2 := &LinkedListNode{}

	for _, v := range values {
		l1.addNode(v)
	}

	values = []int{5, 9, 2, 3}
	for _, v := range values {
		l2.addNode(v)
	}

	l1.sumLists(l2)

	example := "->2->1->9->4"
	if l1.toString() != example {
		t.Errorf("FAIL, the list should be: %s, and got: %s", example, l1.toString())
	}

	fmt.Println("Test Finished")
}
