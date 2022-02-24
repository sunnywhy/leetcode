package main

// Convert BST to Greater Tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var prev int

func convertBST(root *TreeNode) *TreeNode {
	prev = 0
	convert(root)
	return root
}

func convert(root *TreeNode) {
	if root == nil {
		return
	}
	convert(root.Right)
	temp := root.Val
	root.Val += prev
	prev += temp
	convert(root.Left)
}
