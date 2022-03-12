package main

import "container/list"

// Shortest Path in Binary Matrix

func shortestPathBinaryMatrix(grid [][]int) int {
	n := len(grid)

	if grid[0][0] == 1 || grid[n-1][n-1] == 1 {
		return -1
	}

	visited := make([][]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = make([]bool, n)
	}

	q := list.New()
	q.PushBack([]int{0, 0})
	visited[0][0] = true
	var steps int
	dirs := [][]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}, {-1, -1}, {-1, 1}, {1, -1}, {1, 1}}
	for q.Len() > 0 {
		steps++
		size := q.Len()
		for i := 0; i < size; i++ {
			cur := q.Front()
			q.Remove(cur)
			x := cur.Value.([]int)[0]
			y := cur.Value.([]int)[1]
			if x == n-1 && y == n-1 {
				return steps
			}
			for j := 0; j < len(dirs); j++ {
				nextX := x + dirs[j][0]
				nextY := y + dirs[j][1]
				if nextX < 0 || nextX >= n || nextY < 0 || nextY >= n {
					continue
				}
				if grid[nextX][nextY] == 1 || visited[nextX][nextY] {
					continue
				}
				q.PushBack([]int{nextX, nextY})
				visited[nextX][nextY] = true
			}
		}
	}
	return -1
}
