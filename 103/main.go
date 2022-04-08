package main

import "container/list"

// Binary Tree Zigzag Level Order Traversal
// https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	q := list.New()
	q.PushBack(root)
	var leftToRight bool
	for q.Len() > 0 {
		size := q.Len()
		leftToRight = !leftToRight
		level := make([]int, size)
		for i := 0; i < size; i++ {
			e := q.Front()
			q.Remove(e)
			cur := e.Value.(*TreeNode)
			if leftToRight {
				level[i] = cur.Val
			} else {
				level[size-1-i] = cur.Val
			}
			if cur.Left != nil {
				q.PushBack(cur.Left)
			}
			if cur.Right != nil {
				q.PushBack(cur.Right)
			}
		}
		res = append(res, level)
	}
	return res
}
