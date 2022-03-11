package main

import "sort"

// Accounts Merge

type UnionFind struct {
	Parents []int
}

func Constructor(n int) UnionFind {
	parents := make([]int, n)
	for i := 0; i < n; i++ {
		parents[i] = i
	}
	return UnionFind{parents}
}
func (uf *UnionFind) Find(p int) int {
	parent := uf.Parents[p]
	if parent != p {
		parent = uf.Find(parent)
	}
	return parent
}
func (uf *UnionFind) Union(p, q int) {
	rootP := uf.Find(p)
	rootQ := uf.Find(q)
	if rootP != rootQ {
		uf.Parents[rootP] = rootQ
	}
}

// O(AlogA) A is the total number of the emails
func accountsMerge(accounts [][]string) [][]string {
	emailToIds := make(map[string]int)
	idToNames := make(map[int]string)
	id := 0

	for _, v := range accounts {
		name := v[0]
		for i := 1; i < len(v); i++ {
			email := v[i]
			if _, ok := emailToIds[email]; ok {
				continue
			}
			emailToIds[email] = id
			idToNames[id] = name
			id++
		}
	}

	uf := Constructor(len(emailToIds))

	for _, v := range accounts {
		primary := v[1]
		for i := 2; i < len(v); i++ {
			uf.Union(emailToIds[primary], emailToIds[v[i]])
		}
	}

	idToEmailArray := make(map[int][]string)
	for k, v := range emailToIds {
		rootId := uf.Find(v)
		idToEmailArray[rootId] = append(idToEmailArray[rootId], k)
	}

	var res [][]string
	for k, v := range idToEmailArray {
		name := idToNames[k]
		sort.Strings(v)
		res = append(res, append([]string{name}, v...))
	}
	return res
}
