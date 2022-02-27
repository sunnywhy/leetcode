package main

// Convert Binary Search Tree to Sorted Doubly Linked List

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

var head *Node
var cur *Node

func treeToDoublyList(root *Node) *Node {
	if root == nil {
		return root
	}
	head = &Node{}
	cur = head
	inorder(root)
	head.Right.Left = cur
	cur.Right = head.Right
	return head.Right
}
func inorder(root *Node) {
	if root == nil {
		return
	}
	inorder(root.Left)
	cur.Right = root
	root.Left = cur
	cur = root
	inorder(root.Right)
}
