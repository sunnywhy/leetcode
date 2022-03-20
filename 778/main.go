package main

import (
	"container/heap"
	"math"
)

// Swim in Rising Water

type Position struct {
	Row  int
	Col  int
	Time int
}

type minHeap []Position

func (h minHeap) Len() int {
	return len(h)
}
func (h minHeap) Less(i, j int) bool {
	return h[i].Time < h[j].Time
}
func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(Position))
}
func (h *minHeap) Pop() interface{} {
	n := h.Len()
	val := (*h)[n-1]
	*h = (*h)[:n-1]
	return val
}

func swimInWater(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	h := &minHeap{}
	start := Position{0, 0, grid[0][0]}
	heap.Push(h, start)
	visited[0][0] = true

	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for h.Len() > 0 {
		cur := heap.Pop(h).(Position)
		curRow, curCol, curTime := cur.Row, cur.Col, cur.Time
		if curRow == m-1 && curCol == n-1 {
			return curTime
		}
		for i := 0; i < len(dirs); i++ {
			nextRow, nextCol := curRow+dirs[i][0], curCol+dirs[i][1]
			if nextRow < 0 || nextRow >= m || nextCol < 0 || nextCol >= n || visited[nextRow][nextCol] {
				continue
			}
			nextTime := int(math.Max(float64(curTime), float64(grid[nextRow][nextCol])))
			heap.Push(h, Position{nextRow, nextCol, nextTime})
			visited[nextRow][nextCol] = true
		}
	}
	return -1
}
