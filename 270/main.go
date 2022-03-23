package main

import "math"

// Closest Binary Search Tree Value

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var diff float64
var res int

func closestValue(root *TreeNode, target float64) int {
	diff = math.MaxFloat64
	check(root, target)
	return res
}

func check(root *TreeNode, target float64) {
	if root == nil {
		return
	}

	cur := math.Abs(float64(root.Val) - target)
	if diff > cur {
		diff = cur
		res = root.Val
	}
	if float64(root.Val) > target {
		check(root.Left, target)
	} else {
		check(root.Right, target)
	}
}
