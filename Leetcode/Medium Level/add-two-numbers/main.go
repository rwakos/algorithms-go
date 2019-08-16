package main

/*
ListNode bla
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

/*
Runtime: 4 ms, faster than 98.96% of Go online submissions for Add Two Numbers.
Memory Usage: 4.7 MB, less than 90.24% of Go online submissions for Add Two Numbers.
*/
func addTwoNumbers(l *ListNode, l2 *ListNode) *ListNode {
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
			node := &ListNode{Val: 0, Next: nil}
			l1.Next = node
		} else if l1.Next != nil && l2.Next == nil { //L1 is longer
			node := &ListNode{Val: 0, Next: nil}
			l2.Next = node
		}

		if over > 0 && l1.Next == nil && l2.Next == nil { //we have over in the last digit
			node := &ListNode{Val: over, Next: nil}
			l1.Next = node
			return l
		}

		l2 = l2.Next
	}
	return l
}
