package main

// Invert Binary Tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	node := TreeNode{Val: root.Val}
	node.Left = invertTree(root.Right)
	node.Right = invertTree(root.Left)
	return &node
}
