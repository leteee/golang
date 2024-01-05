package sort

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

func ReverseList(head *Node) *Node {
	var prev *Node
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

func TestReverseList() {
	node := &Node{1, &Node{2, &Node{3, &Node{4, &Node{5, nil}}}}}
	node = ReverseList(node)
	for node != nil {
		fmt.Printf("%v \t", node.Val)
		node = node.Next
	}
}
