package main

import "fmt"

/*
LinkedList simple linked list
*/
type LinkedList struct {
	Val  int
	Next *LinkedList
}

func main() {
	l1 := new(LinkedList)
	l2 := new(LinkedList)

	l1.addNode(2)
	l1.addNode(3)
	l1.addNode(4)
	l1.print()

	l2.addNode(4)
	l2.addNode(7)
	l2.addNode(9)
	l2.print()

	l3 := sumTwo(l1, l2)
	l3.print()

}

func sumTwo(l *LinkedList, l2 *LinkedList) *LinkedList {
	over := 0
	sum := 0

	for l1 := l; l1 != nil; l1 = l1.Next {
		sum = l1.Val + l2.Val + over
		over = 0

		if sum%10 >= 0 && sum != sum%10 {
			over++
			sum = sum % 10
		}

		l1.Val = sum

		if l1.Next == nil && l2.Next != nil { //L2 is longer
			node := &LinkedList{Val: 0, Next: nil}
			l1.Next = node
		} else if l1.Next != nil && l2.Next == nil { //L1 is longer
			node := &LinkedList{Val: 0, Next: nil}
			l2.Next = node
		}

		if over > 0 && l1.Next == nil && l2.Next == nil { //we have over in the last digit
			node := &LinkedList{Val: over, Next: nil}
			l1.Next = node
			return l
		}

		l2 = l2.Next
	}
	return l
}

func (l *LinkedList) addNode(val int) {
	if l.Val == 0 && l.Next == nil {
		l.Val = val
		return
	}

	for l1 := l; l1 != nil; l1 = l1.Next {
		if l1.Next == nil {
			node := &LinkedList{Val: val, Next: nil}
			l1.Next = node
			return
		}
	}
}

func (l *LinkedList) print() {
	for l1 := l; l1 != nil; l1 = l1.Next {
		fmt.Print("->", l1.Val)
	}
	fmt.Println("")
}
