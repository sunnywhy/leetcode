package main

// Lowest Common Ancestor of a Binary Tree II
// https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree-ii/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var count int

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	count = 0
	lca := findLCA(root, p, q)
	if count == 2 {
		return lca
	}
	return nil
}
func findLCA(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root == p || root == q {
		count++
		return root
	}
	left := findLCA(root.Left, p, q)
	right := findLCA(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	}
	return right
}

func main() {
	root := &TreeNode{Val: 3}
	left := &TreeNode{Val: 5, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 7}, Right: &TreeNode{Val: 4}}}
	right := &TreeNode{Val: 1, Left: &TreeNode{Val: 0}, Right: &TreeNode{Val: 8}}
	root.Left = left
	root.Right = right
	result := lowestCommonAncestor(root, left, right)
	println(result.Val)
	result = lowestCommonAncestor(root, left, &TreeNode{Val: 108})
	println(result == nil)
}
