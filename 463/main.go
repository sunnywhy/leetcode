package main

// Island Perimeter
// https://leetcode.com/problems/island-perimeter/

func islandPerimeter(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				return dfs(&grid, i, j, &visited)
			}
		}
	}
	return 0
}

var dirs = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func dfs(grid *[][]int, i, j int, visited *[][]bool) int {
	m, n := len(*grid), len((*grid)[0])
	res := 0
	(*visited)[i][j] = true
	for k := 0; k < len(dirs); k++ {
		row, col := i+dirs[k][0], j+dirs[k][1]
		if row < 0 || row >= m || col < 0 || col >= n || (*grid)[row][col] == 0 {
			res++
			continue
		}
		if (*visited)[row][col] {
			continue
		}
		res += dfs(grid, row, col, visited)
	}
	return res
}

// count, only need to check up and left
func islandPerimeter2(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				res += 4
				if i > 0 && grid[i-1][j] == 1 {
					res -= 2
				}
				if j > 0 && grid[i][j-1] == 1 {
					res -= 2
				}
			}
		}
	}
	return res
}
