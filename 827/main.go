package main

// Making A Large Island

var dirs = [][]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}

func largestIsland(grid [][]int) int {
	n := len(grid)
	var res int
	islandId := 1
	idMap := make(map[int]int)  // island id -> size
	visited := make([][]int, n) // position -> id
	for i := 0; i < n; i++ {
		visited[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 || visited[i][j] > 0 {
				continue
			}
			size := dfs(grid, i, j, &visited, islandId)
			idMap[islandId] = size
			if size > res {
				res = size
			}
			islandId++
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				continue
			}
			size := calculate(grid, i, j, &visited, idMap)
			if size > res {
				res = size
			}
		}
	}
	return res
}

func dfs(grid [][]int, r, c int, visited *[][]int, id int) int {
	n := len(grid)
	size := 1
	(*visited)[r][c] = id
	for _, v := range dirs {
		nextRow := r + v[0]
		nextCol := c + v[1]
		if nextRow < 0 || nextRow >= n || nextCol < 0 || nextCol >= n || grid[nextRow][nextCol] == 0 {
			continue
		}
		if (*visited)[nextRow][nextCol] > 0 {
			continue
		}
		size += dfs(grid, nextRow, nextCol, visited, id)
	}
	return size
}

func calculate(grid [][]int, r, c int, visited *[][]int, idMap map[int]int) int {
	n := len(grid)
	size := 1
	checkedIds := make(map[int]bool)
	for _, v := range dirs {
		nextRow := r + v[0]
		nextCol := c + v[1]
		if nextRow < 0 || nextRow >= n || nextCol < 0 || nextCol >= n || grid[nextRow][nextCol] == 0 {
			continue
		}
		id := (*visited)[nextRow][nextCol]
		if checkedIds[id] {
			continue
		}
		size += idMap[id]
		checkedIds[id] = true
	}
	return size
}
