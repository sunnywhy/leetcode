package main

// Delete Nodes And Return Forest
// https://leetcode.com/problems/delete-nodes-and-return-forest/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
	var res []*TreeNode
	toDeleteMap := make(map[int]bool)
	for _, v := range to_delete {
		toDeleteMap[v] = true
	}
	dfs(root, true, toDeleteMap, &res)
	return res
}
func dfs(root *TreeNode, isRoot bool, toDeleteMap map[int]bool, res *[]*TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	isDelete := toDeleteMap[root.Val]

	if !isDelete && isRoot {
		*res = append(*res, root)
	}
	root.Left = dfs(root.Left, isDelete, toDeleteMap, res)
	root.Right = dfs(root.Right, isDelete, toDeleteMap, res)

	if isDelete {
		return nil
	}
	return root
}
