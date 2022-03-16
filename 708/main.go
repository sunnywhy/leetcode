package main

// Insert into a Sorted Circular Linked List

type Node struct {
	Val  int
	Next *Node
}

func insert(aNode *Node, x int) *Node {
	node := &Node{Val: x}
	if aNode == nil {
		node.Next = node
		return node
	}

	prev := aNode
	cur := prev.Next

	for cur != aNode {
		if x >= prev.Val && x <= cur.Val {
			break
		}

		if cur.Val < prev.Val {
			if x >= prev.Val || x <= cur.Val {
				break
			}
		}
		prev = prev.Next
		cur = cur.Next
	}

	prev.Next = node
	node.Next = cur
	return aNode
}
