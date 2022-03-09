package main

import "math"

// Diameter of Binary Tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res int

func diameterOfBinaryTree(root *TreeNode) int {
	res = 0
	calculate(root)
	return res
}
func calculate(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 0
	}
	left := calculate(root.Left)
	right := calculate(root.Right)
	total := 0
	if root.Left != nil {
		total += left + 1
	}
	if root.Right != nil {
		total += right + 1
	}
	if total > res {
		res = total
	}
	return int(math.Max(float64(left), float64(right))) + 1
}
