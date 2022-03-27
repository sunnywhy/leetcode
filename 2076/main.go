package main

// Process Restricted Friend Requests

type UnionFind struct {
	parents []int
}

func Constructor(n int) *UnionFind {
	parents := make([]int, n)
	for i := 0; i < n; i++ {
		parents[i] = i
	}
	return &UnionFind{parents}
}
func (uf *UnionFind) Find(p int) int {
	parent := uf.parents[p]
	if p != parent {
		parent = uf.Find(parent)
	}
	return parent
}
func (uf *UnionFind) Union(p, q int) {
	rootP := uf.Find(p)
	rootQ := uf.Find(q)
	if rootP != rootQ {
		uf.parents[rootP] = rootQ
	}
}
func (uf UnionFind) Connected(p, q int) bool {
	rootP := uf.Find(p)
	rootQ := uf.Find(q)
	return rootP == rootQ
}
func friendRequests(n int, restrictions [][]int, requests [][]int) []bool {
	uf := Constructor(n)
	res := make([]bool, len(requests))
	for i, request := range requests {
		rootP := uf.Find(request[0])
		rootQ := uf.Find(request[1])
		if rootP == rootQ {
			res[i] = true
			continue
		}

		if valid(restrictions, rootP, rootQ, uf) {
			uf.Union(rootP, rootQ)
			res[i] = true
		}
	}
	return res
}

func valid(restrictions [][]int, p, q int, uf *UnionFind) bool {
	for _, restriction := range restrictions {
		root1 := uf.Find(restriction[0])
		root2 := uf.Find(restriction[1])
		if (root1 == p && root2 == q) || (root1 == q && root2 == p) {
			return false
		}
	}
	return true
}
