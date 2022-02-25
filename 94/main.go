package main

// Binary Tree Inorder Traversal
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var res []int
	inorder(root, &res)
	return res
}
func inorder(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	inorder(root.Left, res)
	*res = append(*res, root.Val)
	inorder(root.Right, res)
}

func inorderTraversalIterative(root *TreeNode) []int {
	var res []int
	var stack []*TreeNode
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		n := len(stack)
		cur := stack[n-1]
		stack = stack[:n-1]
		res = append(res, cur.Val)
		root = cur.Right
	}
	return res
}
