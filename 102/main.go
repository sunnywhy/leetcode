package main

import "container/list"

// Binary Tree Level Order Traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	q := list.New()
	q.PushBack(root)
	for q.Len() > 0 {
		size := q.Len()
		var level []int
		for i := 0; i < size; i++ {
			cur := q.Front()
			node := cur.Value.(*TreeNode)
			q.Remove(cur)
			level = append(level, node.Val)
			if node.Left != nil {
				q.PushBack(node.Left)
			}
			if node.Right != nil {
				q.PushBack(node.Right)
			}
		}
		res = append(res, level)
	}
	return res
}
