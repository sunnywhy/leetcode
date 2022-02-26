package main

import "math"

// Maximum Depth of Binary Tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	return calculateDepth(root)
}
func calculateDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := calculateDepth(root.Left)
	right := calculateDepth(root.Right)
	return int(math.Max(float64(left), float64(right))) + 1
}
