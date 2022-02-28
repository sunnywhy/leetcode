package main

import "math"

// Binary Tree Maximum Path Sum

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res int

func maxPathSum(root *TreeNode) int {
	res = math.MinInt
	calculate(root)
	return res
}
func calculate(root *TreeNode) int {
	if root == nil {
		return 0
	}
	total := root.Val
	left := calculate(root.Left)
	right := calculate(root.Right)

	if left > 0 {
		total += left
	}
	if right > 0 {
		total += right
	}
	if total > res {
		res = total
	}
	if left > 0 || right > 0 {
		return root.Val + int(math.Max(float64(left), float64(right)))
	} else {
		return root.Val
	}
}
