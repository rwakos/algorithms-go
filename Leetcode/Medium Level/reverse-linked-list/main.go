package main

import "fmt"

func main() {
	l1 := new(ListNode)
	l1.addNode(1)
	l1.addNode(2)
	l1.addNode(3)
	l1.addNode(4)
	l1.addNode(5)

	l2 := reverseBetween(l1, 2, 4)
	fmt.Println(l2.toString())
}

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Reverse Linked List II.
Memory Usage: 2 MB, less than 75.00% of Go online submissions for Reverse Linked List II.
 * ListNode
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if m == n || m > n {
		return head
	}
	startChange := false
	j := n
	l := head
	for i := 1; i <= n; i++ {
		if i == m {
			startChange = true
		}

		if startChange {
			l2 := l
			for k := i; k <= j; k++ {
				if k == j {
					temp := l2.Val
					l2.Val = l.Val
					l.Val = temp
					j--
				}
				l2 = l2.Next
			}
		}

		l = l.Next
	}
	return head
}

func (l *ListNode) toString() string {
	s := ""
	for l1 := l; l1 != nil; l1 = l1.Next {
		s = fmt.Sprintf("%s->%d", s, l1.Val)
	}
	return s
}

func (l *ListNode) addNode(val int) {
	if l == nil || (l.Val == 0 && l.Next == nil) {
		*l = *&ListNode{Val: val, Next: nil}
		return
	}

	for l1 := l; l1 != nil; l1 = l1.Next {
		if l1.Next == nil {
			node := &ListNode{val, nil}
			l1.Next = node
			return
		}
	}
}
