package main

import (
	"container/list"
	"math"
)

// Find Largest Value in Each Tree Row

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func largestValues(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	q := list.New()
	q.PushBack(root)
	for q.Len() > 0 {
		size := q.Len()
		max := math.MinInt
		for i := 0; i < size; i++ {
			e := q.Front()
			q.Remove(e)
			cur := e.Value.(*TreeNode)
			if cur.Val > max {
				max = cur.Val
			}
			if cur.Left != nil {
				q.PushBack(cur.Left)
			}
			if cur.Right != nil {
				q.PushBack(cur.Right)
			}
		}
		res = append(res, max)
	}
	return res
}
