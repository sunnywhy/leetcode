package main

import "math"

// Leftmost Column with at Least a One

type BinaryMatrix struct {
	Get        func(int, int) int
	Dimensions func() []int
}

// O(M*logN)
func leftMostColumnWithOne(binaryMatrix BinaryMatrix) int {
	dimensions := binaryMatrix.Dimensions()
	m, n := dimensions[0], dimensions[1]
	res := math.MaxInt
	for i := 0; i < m; i++ {
		idx := getIndex(&binaryMatrix, i, n)
		if res > idx {
			res = idx
		}
	}
	if res == math.MaxInt {
		return -1
	}
	return res
}

func getIndex(binaryMatrix *BinaryMatrix, row, n int) int {
	l, r := 0, n-1
	for l <= r {
		m := l + (r-l)/2
		if (*binaryMatrix).Get(row, m) == 1 {
			if m == 0 || (*binaryMatrix).Get(row, m-1) == 0 {
				return m
			} else {
				r = m - 1
			}
		} else {
			l = m + 1
		}
	}
	return math.MaxInt
}

// Method II, O(M + N)
func leftMostColumnWithOne2(binaryMatrix BinaryMatrix) int {
	dimensions := binaryMatrix.Dimensions()
	m, n := dimensions[0], dimensions[1]
	res := -1
	i, j := 0, n-1
	for i < m && j >= 0 {
		if binaryMatrix.Get(i, j) == 1 {
			res = j
			j--
		} else {
			i++
		}
	}
	return res
}
