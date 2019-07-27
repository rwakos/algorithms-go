package main

import (
	"fmt"
)

//NodeList some stuff
type NodeList struct {
	Val  int
	Next *NodeList
}

func main() {
	l := NodeList{
		Val: 1,
	}

	l.pushNode(2)
	l.pushNode(3)

	fmt.Println(l.Val)
	fmt.Println(l.Next.Val)
	fmt.Println(l.Next.Next.Val)
}

func (n *NodeList) pushNode(v int) {
	s := &NodeList{
		Val:  v,
		Next: nil,
	}

	if n == nil {
		n = s
	} else {
		currentNode := n
		for currentNode.Next != nil {
			currentNode = currentNode.Next
		}
		currentNode.Next = s
	}
}
