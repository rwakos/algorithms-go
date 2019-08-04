package main

import "fmt"

/*
LinkedListNode simple linked list
*/
type LinkedListNode struct {
	data int
	next *LinkedListNode
}

func main() {
	t := []int{1, 2, 3, 4, 5, 2, 1, 3, 4, 5}
	l := new(LinkedListNode)
	l = nil
	for _, v := range t {
		item := &LinkedListNode{v, nil}
		l = addNode(item, l)
	}
	fmt.Println("Original: ")
	printLinkedList(l)
	l = removeDuplicates(l)
	fmt.Println("After Removal: ")
	printLinkedList(l)
}

func removeDuplicates(lln *LinkedListNode) *LinkedListNode {
	l := lln

	for l != nil {
		runner := l
		for runner.next != nil {
			if runner.next.data == l.data {
				runner.next = runner.next.next
			} else {
				runner = runner.next
			}
		}
		l = l.next
	}

	return lln
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
