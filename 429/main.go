package main

import "container/list"

// N-ary Tree Level Order Traversal

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
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
			q.Remove(cur)
			node := cur.Value.(*Node)
			level = append(level, node.Val)
			for _, child := range node.Children {
				if child == nil {
					continue
				}
				q.PushBack(child)
			}
		}
		res = append(res, level)
	}
	return res
}
