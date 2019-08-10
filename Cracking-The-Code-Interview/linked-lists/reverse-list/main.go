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
}

func (l *LinkedListNode) reverseList() {
	if l == nil {
		return
	}
	reversed := &LinkedListNode{data: l.data, next: nil}
	for l1 := l.next; l1 != nil; l1 = l1.next {
		node := &LinkedListNode{data: l1.data, next: reversed}
		reversed = node
	}
	*l = *reversed
}

func (l *LinkedListNode) toString() string {
	s := ""
	for l1 := l; l1 != nil; l1 = l1.next {
		s = fmt.Sprintf("%s->%d", s, l1.data)
	}
	return s
}

func (l *LinkedListNode) addNode(val int) {
	if l == nil || (l.data == 0 && l.next == nil) {
		*l = *&LinkedListNode{data: val, next: nil}
		return
	}

	for l1 := l; l1 != nil; l1 = l1.next {
		if l1.next == nil {
			node := &LinkedListNode{val, nil}
			l1.next = node
			return
		}
	}
}
