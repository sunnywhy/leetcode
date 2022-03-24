package main

// Number of Valid Move Combinations On Chessboard

type Point struct {
	x int
	y int
}

var res int

func countCombinations(pieces []string, positions [][]int) int {
	n := len(pieces)

	targets := make([][]Point, n)

	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}, {-1, -1}, {-1, 1}, {1, -1}, {1, 1}}
	for i, p := range pieces {
		curX, curY := positions[i][0], positions[i][1]
		targets[i] = append(targets[i], Point{curX, curY})

		for j := 0; j < len(dirs); j++ {
			if p == "rook" && j >= 4 {
				continue
			} else if p == "bishop" && j < 4 {
				continue
			}
			nextX, nextY := curX+dirs[j][0], curY+dirs[j][1]

			for nextX >= 1 && nextX <= 8 && nextY >= 1 && nextY <= 8 {
				targets[i] = append(targets[i], Point{nextX, nextY})
				nextX += dirs[j][0]
				nextY += dirs[j][1]
			}
		}
	}

	res = 0
	var curTarget []Point
	dfs(&positions, &targets, n, 0, &curTarget)
	return res
}

func dfs(positions *[][]int, targets *[][]Point, n, idx int, curTarget *[]Point) {
	if idx == n {
		if valid(positions, curTarget, n) {
			res++
		}
		return
	}
	for i := 0; i < len((*targets)[idx]); i++ {
		*curTarget = append(*curTarget, (*targets)[idx][i])
		dfs(positions, targets, n, idx+1, curTarget)
		*curTarget = (*curTarget)[:len(*curTarget)-1]
	}
}

func valid(positions *[][]int, curTarget *[]Point, n int) bool {
	copies := make([][]int, n)
	for i := 0; i < n; i++ {
		copies[i] = []int{(*positions)[i][0], (*positions)[i][1]}
	}
	keepMoving := true
	for keepMoving {
		keepMoving = false
		used := make(map[int]bool)
		for i := 0; i < n; i++ {
			curX, curY := copies[i][0], copies[i][1]
			targetX, targetY := (*curTarget)[i].x, (*curTarget)[i].y
			var diffX, diffY int
			if targetX > curX {
				diffX = 1
			} else if targetX < curX {
				diffX = -1
			}
			if targetY > curY {
				diffY = 1
			} else if targetY < curY {
				diffY = -1
			}
			if diffX != 0 || diffY != 0 {
				keepMoving = true
				copies[i][0] += diffX
				copies[i][1] += diffY
			}
			key := 13*copies[i][0] + copies[i][1]
			if used[key] {
				return false
			}
			used[key] = true
		}
	}
	return true
}
