package main

import "math"

// Balanced Binary Tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var balanced bool

func isBalanced(root *TreeNode) bool {
	balanced = true
	checkHeight(root)
	return balanced
}
func checkHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if !balanced {
		return -1
	}
	leftHeight := checkHeight(root.Left)
	rightHeight := checkHeight(root.Right)
	if math.Abs(float64(leftHeight)-float64(rightHeight)) > 1 {
		balanced = false
	}
	return int(math.Max(float64(leftHeight), float64(rightHeight))) + 1
}
