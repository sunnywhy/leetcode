package main

// Diagonal Traverse

// for a specific diagonal, i+j always is the same value
// Time: O(M * N)  Space: O(M * N)
func findDiagonalOrder(mat [][]int) []int {
	m, n := len(mat), len(mat[0])
	data := make([][]int, m+n+1)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			data[i+j] = append(data[i+j], mat[i][j])
		}
	}

	var res []int
	for k, arr := range data {
		if k%2 == 1 {
			for _, v := range arr {
				res = append(res, v)
			}
		} else {
			for i := len(arr) - 1; i >= 0; i-- {
				res = append(res, arr[i])
			}
		}
	}
	return res
}

// simulation
func findDiagonalOrder2(mat [][]int) []int {
	m, n := len(mat), len(mat[0])
	var row, col, nextRow, nextCol int
	var res []int
	upward := true

	for row < m && col < n {
		res = append(res, mat[row][col])
		if upward {
			nextRow, nextCol = row-1, col+1
		} else {
			nextRow, nextCol = row+1, col-1
		}

		if nextRow < 0 || nextRow >= m || nextCol < 0 || nextCol >= n {
			if upward {
				if col == n-1 {
					nextRow, nextCol = row+1, col
				} else {
					nextRow, nextCol = row, col+1
				}
			} else {
				if row == m-1 {
					nextRow, nextCol = row, col+1
				} else {
					nextRow, nextCol = row+1, col
				}
			}
			upward = !upward
		}

		row, col = nextRow, nextCol
	}
	return res
}
