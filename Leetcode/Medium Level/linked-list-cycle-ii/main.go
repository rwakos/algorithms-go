package main

/**
https://leetcode.com/problems/linked-list-cycle-ii/submissions/
Find the node where the cycle starts

Runtime: 8 ms, faster than 41.12% of Go online submissions for Linked List Cycle II.
Memory Usage: 5 MB, less than 16.67% of Go online submissions for Linked List Cycle II.
*/

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func detectCycle(head *ListNode) *ListNode {
	myMap := make(map[*ListNode]bool)
	myMap[head] = true
	for l := head; l != nil; l = l.Next {
		_, ok := myMap[l.Next]
		if ok {
			return l.Next
		}
		myMap[l.Next] = true
	}

	return nil
}
