package main

// Smallest Subtree with all the Deepest Nodes
// https://leetcode.com/problems/smallest-subtree-with-all-the-deepest-nodes/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func subtreeWithAllDeepest(root *TreeNode) *TreeNode {
	_, node := dfs(root)
	return node
}

func dfs(root *TreeNode) (int, *TreeNode) {
	if root == nil {
		return 0, nil
	}
	leftDepth, left := dfs(root.Left)
	rightDepth, right := dfs(root.Right)
	if leftDepth > rightDepth {
		return leftDepth + 1, left
	} else if leftDepth < rightDepth {
		return rightDepth + 1, right
	}
	return leftDepth + 1, root
}
