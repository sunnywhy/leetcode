package main

import (
	"container/list"
	"math"
)

// Shortest Distance from All Buildings

// O(M^2*N^2)
func shortestDistance(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	reaches, distances := make([][]int, m), make([][]int, m)
	for i := 0; i < m; i++ {
		reaches[i] = make([]int, n)
		distances[i] = make([]int, n)
	}
	var buildings int

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				bfs(grid, i, j, &reaches, &distances)
				buildings++
			}
		}
	}

	res := math.MaxInt
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if reaches[i][j] == buildings && distances[i][j] < res {
				res = distances[i][j]
			}
		}
	}

	if res == math.MaxInt {
		return -1
	}
	return res
}

func bfs(grid [][]int, i, j int, reaches, distances *[][]int) {
	m, n := len(grid), len(grid[0])
	dirs := [][]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}
	q := list.New()
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	q.PushBack([]int{i, j})
	visited[i][j] = true
	var steps int
	for q.Len() > 0 {
		steps++
		size := q.Len()
		for i := 0; i < size; i++ {
			cur := q.Front()
			q.Remove(cur)
			x := cur.Value.([]int)[0]
			y := cur.Value.([]int)[1]

			for k := 0; k < len(dirs); k++ {
				nextX, nextY := x+dirs[k][0], y+dirs[k][1]
				if nextX < 0 || nextX >= m || nextY < 0 || nextY >= n || grid[nextX][nextY] != 0 {
					continue
				}
				if visited[nextX][nextY] {
					continue
				}
				(*reaches)[nextX][nextY]++
				(*distances)[nextX][nextY] += steps
				q.PushBack([]int{nextX, nextY})
				visited[nextX][nextY] = true
			}
		}
	}
}
