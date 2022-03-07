package main

import "sort"

// Vertical Order Traversal of a Binary Tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type pair struct {
	Row int
	Val int
}

var maxCol, minCol int

// traversal O(N)
// sort, assume the average node in each col is M, the total node is N
// then sort needs N/ M * MLogM = N * LogM
func verticalTraversal(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	maxCol = 0
	minCol = 0
	m := make(map[int][]*pair)
	traversal(root, m, 0, 0)
	for i := minCol; i <= maxCol; i++ {
		cur := m[i]
		sort.Slice(cur, func(i, j int) bool {
			if cur[i].Row == cur[j].Row {
				return cur[i].Val < cur[j].Val
			} else {
				return cur[i].Row < cur[j].Row
			}
		})
		var levels []int
		for _, v := range cur {
			levels = append(levels, v.Val)
		}
		res = append(res, levels)
	}
	return res
}
func traversal(root *TreeNode, m map[int][]*pair, row, col int) {
	if root == nil {
		return
	}
	m[col] = append(m[col], &pair{row, root.Val})
	if col > maxCol {
		maxCol = col
	}
	if col < minCol {
		minCol = col
	}
	traversal(root.Left, m, row+1, col-1)
	traversal(root.Right, m, row+1, col+1)
}
