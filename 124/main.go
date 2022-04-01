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
	checkAndCalculate(root)
	if res == math.MinInt {
		return 0
	}
	return res
}

func checkAndCalculate(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := checkAndCalculate(root.Left)
	right := checkAndCalculate(root.Right)
	if left < 0 {
		left = 0
	}
	if right < 0 {
		right = 0
	}
	if res < left+right+root.Val {
		res = left + right + root.Val
	}
	return root.Val + int(math.Max(float64(left), float64(right)))
}
