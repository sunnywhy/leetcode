package main

// Sum Root to Leaf Numbers

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	var total int
	dfs(root, 0, &total)
	return total
}

func dfs(root *TreeNode, cur int, total *int) {
	if root == nil {
		return
	}

	cur = cur*10 + root.Val
	if root.Left == nil && root.Right == nil {
		*total += cur
		return
	}
	dfs(root.Left, cur, total)
	dfs(root.Right, cur, total)
}
