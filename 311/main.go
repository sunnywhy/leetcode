package main

// Sparse Matrix Multiplication
// https://leetcode.com/problems/sparse-matrix-multiplication/
func multiply(mat1 [][]int, mat2 [][]int) [][]int {
	m, k, n := len(mat1), len(mat2), len(mat2[0])
	res := make([][]int, m)
	for i := 0; i < m; i++ {
		res[i] = make([]int, n)
	}

	data1, data2 := make(map[[2]int]int), make(map[[2]int]int)
	for i := 0; i < m; i++ {
		for j := 0; j < k; j++ {
			if mat1[i][j] != 0 {
				data1[[2]int{i, j}] = mat1[i][j]
			}
		}
	}

	for i := 0; i < k; i++ {
		for j := 0; j < n; j++ {
			if mat2[i][j] != 0 {
				data2[[2]int{i, j}] = mat2[i][j]
			}
		}
	}

	for k1, v1 := range data1 {
		for k2, v2 := range data2 {
			if k1[1] == k2[0] {
				res[k1[0]][k2[1]] += v1 * v2
			}
		}
	}
	return res
}
