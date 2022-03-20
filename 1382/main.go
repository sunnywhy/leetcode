package main

// Balance a Binary Search Tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// O(N)
func balanceBST(root *TreeNode) *TreeNode {
	var values []int
	traversal(root, &values)
	return buildTree(&values, 0, len(values)-1)
}

func traversal(root *TreeNode, values *[]int) {
	if root == nil {
		return
	}
	traversal(root.Left, values)
	*values = append(*values, root.Val)
	traversal(root.Right, values)
}

func buildTree(values *[]int, start, end int) *TreeNode {
	if start > end {
		return nil
	}
	m := start + (end-start)/2
	node := &TreeNode{Val: (*values)[m]}
	node.Left = buildTree(values, start, m-1)
	node.Right = buildTree(values, m+1, end)
	return node
}
