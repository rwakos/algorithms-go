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
	t := []int{1, 2, 3, 4, 5}
	l := new(LinkedListNode)
	l = nil
	for _, v := range t {
		item := &LinkedListNode{v, nil}
		l = addNode(item, l)
	}
	fmt.Println("Original: ")
	printLinkedList(l)
	l = deleteMiddleNode(l)
	fmt.Println("After Removal: ")
	printLinkedList(l)
}

func deleteMiddleNode(lln *LinkedListNode) *LinkedListNode {
	count := findMiddle(lln)

	if count < 3 {
		return lln
	}

	middle := count / 2
	if count%2 > 0 {
		middle++
	}
	l := lln
	i := 1
	for l.next != nil {
		if i == middle {
			l.next = l.next.next
			return lln
		}
		l = l.next
		i++
	}
	return lln
}

func findMiddle(lln *LinkedListNode) int {
	middle := 0
	for l := lln; l.next != nil; l = l.next {
		middle++
	}
	return middle
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
