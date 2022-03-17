package main

import "container/list"

// Clone Graph

type Node struct {
	Val       int
	Neighbors []*Node
}

// DFS O(V + E)
func cloneGraphDfs(node *Node) *Node {
	if node == nil {
		return node
	}
	visited := make(map[*Node]*Node)
	dfs(node, visited)
	return visited[node]
}

func dfs(node *Node, visited map[*Node]*Node) {
	var newNeighbors []*Node
	newNode := &Node{Val: node.Val}
	visited[node] = newNode
	for _, n := range node.Neighbors {
		if _, ok := visited[n]; !ok {
			dfs(n, visited)
		}
		newNeighbors = append(newNeighbors, visited[n])
	}
	newNode.Neighbors = newNeighbors
}

// BFS
func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	visited := make(map[*Node]*Node)
	q := list.New()
	q.PushBack(node)
	visited[node] = &Node{Val: node.Val}

	for q.Len() > 0 {
		element := q.Front()
		q.Remove(element)
		cur := element.Value.(*Node)
		var newNeighborList []*Node
		for _, neighbor := range cur.Neighbors {
			if _, ok := visited[neighbor]; !ok {
				copyNeighbor := &Node{Val: neighbor.Val}
				q.PushBack(neighbor)
				visited[neighbor] = copyNeighbor
			}
			newNeighborList = append(newNeighborList, visited[neighbor])
		}
		visited[cur].Neighbors = newNeighborList
	}

	return visited[node]
}
