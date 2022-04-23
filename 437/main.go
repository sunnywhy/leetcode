package main

// Path Sum III

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	return pathWithRoot(root, targetSum) + pathSum(root.Left, targetSum) + pathSum(root.Right, targetSum)
}

func pathWithRoot(root *TreeNode, target int) int {
	if root == nil {
		return 0
	}
	total := 0
	target -= root.Val
	if target == 0 {
		total++
	}
	return total + pathWithRoot(root.Left, target) + pathWithRoot(root.Right, target)
}
