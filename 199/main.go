package main

import "container/list"

// Binary Tree Right Side View

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	q := list.New()
	q.PushBack(root)
	for q.Len() > 0 {
		size := q.Len()
		for i := 0; i < size; i++ {
			cur := q.Front()
			curNode := cur.Value.(*TreeNode)
			q.Remove(cur)
			if i == size-1 {
				res = append(res, curNode.Val)
			}
			if curNode.Left != nil {
				q.PushBack(curNode.Left)
			}
			if curNode.Right != nil {
				q.PushBack(curNode.Right)
			}
		}
	}
	return res
}
