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

func reverseList(pList *LinkedListNode) *LinkedListNode {
	pCurr := pList
	var pTop *LinkedListNode
	for {
		if pCurr == nil {
			break
		}
		pTemp := pCurr.next
		pCurr.next = pTop
		pTop = pCurr
		pCurr = pTemp
	}

	return pTop
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
		l.data = val
		return
	}

	for l1 := l; l1 != nil; l1 = l1.next {
		if l1.next == nil {
			node := &LinkedListNode{
				val,
				nil,
			}
			l1.next = node
			return
		}
	}
}
