package main

import (
	"fmt"
)

/*
LinkedListNode basic Linked List
*/
type LinkedListNode struct {
	data int
	next *LinkedListNode
}

func main() {
	values := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	l1 := new(LinkedListNode)
	node := new(LinkedListNode)
	for x := range values {
		l1.addNode(x)
	}

	l2 := l1
	fmt.Println(isCircular(l2)) //False

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

	fmt.Println(isCircular(l1)) //True
}

func isCircular(l1 *LinkedListNode) bool {
	l2 := l1
	for l := l1; l != nil; l = l.next {
		if l2 != nil {
			l2 = l2.next
		}
		if l2 != nil {
			l2 = l2.next
		}
		if l == l2 {
			return true
		}
	}
	return false
}

func (l *LinkedListNode) toString() string {
	s := ""
	i := 0
	for l1 := l; l1 != nil; l1 = l1.next {
		s = fmt.Sprintf("%s->%d", s, l1.data)

		if i > 10 {
			break
		}
		i++
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
