package main

/*
Sum the integer values of two lists, that are stored in reverse mode, and return it
*/
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
}

func (l *LinkedListNode) sumLists(lToAdd *LinkedListNode) {
	over := 0
	l2 := lToAdd
	l1 := l
	var sum int
	//for l1 := l; l1 != nil; l1 = l1.next {
	for l2 = lToAdd; l2 != nil; l2 = l2.next {
		if l1 != nil {
			sum = (l1.data + l2.data + over)
		} else {
			sum = (l2.data + over)
		}

		if sum%10 > 0 {
			if l1 == nil {
				fmt.Println("add", sum)
				l1.addNode(sum % 10)
			} else {
				l1.data = sum % 10
			}
			over = 1
		} else {
			if l1 == nil {
				fmt.Println("add", sum)
				l1.addNode(sum)
			} else {
				l1.data = sum
			}

			over = 0
		}

		if l1 != nil {
			l1 = l1.next
		}
	}

	if over > 0 {
		l1.addNode(over)
	}
}

func (l *LinkedListNode) toString() string {
	s := ""
	for l1 := l; l1 != nil; l1 = l1.next {
		s = fmt.Sprintf("%s->%d", s, l1.data)
	}
	return s
}

func (l *LinkedListNode) addNode(val int) {
	if l == nil {
		node := &LinkedListNode{
			val,
			nil,
		}
		l = node
		return
	}
	if l.data == 0 {
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
