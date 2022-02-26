package main

// Lowest Common Ancestor of a Binary Tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res *TreeNode

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	res = nil
	checkChildrenNumber(root, p, q)
	return res
}

func checkChildrenNumber(root, p, q *TreeNode) int {
	if root == nil {
		return 0
	}
	if res != nil {
		return 2
	}
	rootContains := 0
	if root == p || root == q {
		rootContains = 1
	}
	left := checkChildrenNumber(root.Left, p, q)
	right := checkChildrenNumber(root.Right, p, q)
	if rootContains == 1 && left+right == 1 {
		res = root
		return 2
	} else if left == 1 && right == 1 {
		res = root
		return 2
	}
	return rootContains + left + right
}
