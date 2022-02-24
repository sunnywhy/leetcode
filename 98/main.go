package main

import "math"

// Validate Binary Search Tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var isValid bool

func isValidBST(root *TreeNode) bool {
	isValid = true
	check(root, math.MinInt, math.MaxInt)
	return isValid
}
func check(root *TreeNode, low, high int) bool {
	if isValid == false {
		return false
	}
	if root == nil {
		return true
	}
	if root.Val <= low || root.Val >= high {
		isValid = false
		return false
	}
	return check(root.Left, low, root.Val) && check(root.Right, root.Val, high)
}
