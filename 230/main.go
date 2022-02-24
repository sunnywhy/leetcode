package main

// Kth Smallest Element in a BST

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var count int
var res *TreeNode

func kthSmallest(root *TreeNode, k int) int {
	count = 0
	res = nil
	find(root, k)
	return res.Val
}

func find(root *TreeNode, k int) {
	if root == nil {
		return
	}
	if count > k {
		return
	}
	find(root.Left, k)
	count++
	if count == k {
		res = root
		return
	}
	find(root.Right, k)
}
