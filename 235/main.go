package main

// Lowest Common Ancestor of a Binary Search Tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	return checkAndFind(root, p, q)
}
func checkAndFind(root, p, q *TreeNode) *TreeNode {
	if p.Val < root.Val && q.Val < root.Val {
		return checkAndFind(root.Left, p, q)
	} else if p.Val > root.Val && q.Val > root.Val {
		return checkAndFind(root.Right, p, q)
	} else {
		return root
	}
}
