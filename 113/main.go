package main

// Path Sum II

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res [][]int

func pathSum(root *TreeNode, targetSum int) [][]int {
	res = [][]int{}
	cur := []int{}
	backtracking(root, targetSum, &cur)
	return res
}
func backtracking(root *TreeNode, target int, cur *[]int) {
	if root == nil {
		return
	}
	*cur = append(*cur, root.Val)
	target -= root.Val
	if root.Left == nil && root.Right == nil {
		if target == 0 {
			res = append(res, append([]int{}, *cur...))
		}
	}
	backtracking(root.Left, target, cur)
	backtracking(root.Right, target, cur)
	*cur = (*cur)[:len(*cur)-1]
}
