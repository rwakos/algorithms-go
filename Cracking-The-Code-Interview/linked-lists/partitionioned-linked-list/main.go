package main

import (
	"fmt"
)

/*
LinkedListNode simple linked list
*/
type LinkedListNode struct {
	data int
	next *LinkedListNode
}

func main() {
	t := []int{1, 11, 3, 41, 5, 6, 34, 14, 5, 7, 9}
	l := new(LinkedListNode)
	l = nil
	for _, v := range t {
		item := &LinkedListNode{v, nil}
		l = addNode(item, l)
	}
	fmt.Println("Original: ")
	printLinkedList(l)
	l = partitionLinkedList(8, l)
	fmt.Println("After Partition: ")
	printLinkedList(l)
}

func partitionLinkedList(x int, lln *LinkedListNode) *LinkedListNode {
	partitionedLeft := new(LinkedListNode)
	partitionedRight := new(LinkedListNode)
	lastAddedLeft := new(LinkedListNode)
	partitionedLeft = nil
	partitionedRight = nil
	lastAddedLeft = nil

	for l := lln; l != nil; l = l.next {
		if l.data < x { //Push left
			partitionedLeft = addNodeLeft(&LinkedListNode{l.data, nil}, partitionedLeft)
			lastAddedLeft = l
		} else { //push right
			partitionedRight = addNode(&LinkedListNode{l.data, nil}, partitionedRight)
		}
	}

	lastAddedLeft.next = partitionedRight
	partitionedLeft = addNode(lastAddedLeft, partitionedLeft)

	return partitionedLeft
}

func printLinkedList(lln *LinkedListNode) {
	for l := lln; l != nil; l = l.next {
		fmt.Print("->", l.data)
	}
	fmt.Println("")
}

func addNode(newItem, lln *LinkedListNode) *LinkedListNode {
	if lln == nil {
		return newItem
	}
	for l := lln; l != nil; l = l.next {
		if l.next == nil {
			l.next = newItem
			return lln
		}
	}
	return lln
}

func addNodeLeft(newItem, lln *LinkedListNode) *LinkedListNode {
	if lln == nil {
		return newItem
	}

	left := newItem
	left.next = lln
	return left
}
