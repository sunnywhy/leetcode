package main

import "container/list"

// Check Completeness of a Binary Tree
// https://leetcode.com/problems/check-completeness-of-a-binary-tree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isCompleteTree(root *TreeNode) bool {
	if root == nil {
		return true
	}

	q := list.New()
	q.PushBack(root)
	possibleEnd := false
	for q.Len() > 0 {
		e := q.Front()
		q.Remove(e)
		cur := e.Value.(*TreeNode)

		if possibleEnd && cur != nil {
			return false
		}
		if cur == nil {
			possibleEnd = true
		} else {
			q.PushBack(cur.Left)
			q.PushBack(cur.Right)
		}
	}
	return true
}
