package main

import (
	"container/list"
	"math"
)

// Walls and Gates

//O(M*N)
func wallsAndGates(rooms [][]int) {
	m, n := len(rooms), len(rooms[0])
	q := list.New()

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if rooms[i][j] == 0 {
				q.PushBack([]int{i, j})
			}
		}
	}

	var steps int
	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for q.Len() > 0 {
		steps++
		size := q.Len()
		for i := 0; i < size; i++ {
			e := q.Front()
			q.Remove(e)
			cur := e.Value.([]int)
			for j := 0; j < len(dirs); j++ {
				nextX, nextY := cur[0]+dirs[j][0], cur[1]+dirs[j][1]
				if nextX < 0 || nextX >= m || nextY < 0 || nextY >= n || rooms[nextX][nextY] != math.MaxInt32 {
					continue
				}
				rooms[nextX][nextY] = steps
				q.PushBack([]int{nextX, nextY})
			}
		}
	}
}
