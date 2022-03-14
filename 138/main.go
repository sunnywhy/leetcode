package main

// Copy List with Random Pointer

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	dummy := &Node{}
	tail := dummy
	original := head
	m := make(map[*Node]*Node) // original -> copy

	for original != nil {
		cur := &Node{Val: original.Val}
		m[original] = cur
		tail.Next = cur
		tail = tail.Next
		original = original.Next
	}

	for head != nil {
		m[head].Random = m[head.Random]
		head = head.Next
	}
	return dummy.Next
}
