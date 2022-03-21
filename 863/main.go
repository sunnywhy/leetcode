package main

import "container/list"

// All Nodes Distance K in Binary Tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	graph := make(map[*TreeNode][]*TreeNode)
	buildGraph(graph, root, nil)

	q := list.New()
	visited := make(map[*TreeNode]bool)
	q.PushBack(target)
	visited[target] = true

	var steps int
	var res []int
	for q.Len() > 0 {
		size := q.Len()
		for i := 0; i < size; i++ {
			e := q.Front()
			q.Remove(e)
			cur := e.Value.(*TreeNode)
			if steps == k {
				res = append(res, cur.Val)
			} else {
				for _, neighbor := range graph[cur] {
					if !visited[neighbor] {
						visited[neighbor] = true
						q.PushBack(neighbor)
					}
				}
			}
		}
		steps++
	}
	return res
}

func buildGraph(graph map[*TreeNode][]*TreeNode, node, parent *TreeNode) {
	if node == nil {
		return
	}
	if parent != nil {
		graph[parent] = append(graph[parent], node)
		graph[node] = append(graph[node], parent)
	}
	buildGraph(graph, node.Left, node)
	buildGraph(graph, node.Right, node)
}
