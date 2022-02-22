package main

import "container/list"

// Find Bottom Left Tree Value

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findBottomLeftValue(root *TreeNode) int {
	q := list.New()
	q.PushBack(root)
	var res int
	for q.Len() > 0 {
		size := q.Len()
		for i := 0; i < size; i++ {
			cur := q.Front()
			q.Remove(cur)
			node := cur.Value.(*TreeNode)
			if i == 0 {
				res = node.Val
			}
			if node.Left != nil {
				q.PushBack(node.Left)
			}
			if node.Right != nil {
				q.PushBack(node.Right)
			}
		}
	}
	return res
}
