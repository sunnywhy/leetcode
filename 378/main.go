package main

import (
	"container/heap"
	"math"
)

// Kth Smallest Element in a Sorted Matrix
// https://leetcode.com/problems/kth-smallest-element-in-a-sorted-matrix/

type Node struct {
	row int
	col int
	val int
}
type MinHeap []Node

func (h MinHeap) Len() int {
	return len(h)
}
func (h MinHeap) Less(i, j int) bool {
	return h[i].val < h[j].val
}
func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Node))
}
func (h *MinHeap) Pop() interface{} {
	n := h.Len()
	val := (*h)[n-1]
	*h = (*h)[:n-1]
	return val
}
func kthSmallest(matrix [][]int, k int) int {
	h := &MinHeap{}
	n := len(matrix)
	for i := 0; i < int(math.Min(float64(n), float64(k))); i++ {
		heap.Push(h, Node{i, 0, matrix[i][0]})
	}
	for k-1 > 0 {
		cur := heap.Pop(h).(Node)
		if cur.col < n-1 {
			heap.Push(h, Node{cur.row, cur.col + 1, matrix[cur.row][cur.col+1]})
		}
		k--
	}
	return (*h)[0].val
}
