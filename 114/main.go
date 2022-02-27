package main

// Flatten Binary Tree to Linked List

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var cur *TreeNode

func flatten(root *TreeNode) {
	cur = &TreeNode{}
	preorder(root)
}
func preorder(root *TreeNode) {
	if root == nil {
		return
	}
	left := root.Left
	right := root.Right
	cur.Right = root
	cur = cur.Right
	root.Left = nil
	preorder(left)
	preorder(right)
}
