package main

import "math"

// Product of Two Run-Length Encoded Arrays

func findRLEArray(encoded1 [][]int, encoded2 [][]int) [][]int {
	var res [][]int
	m, n := len(encoded1), len(encoded2)

	for i, j := 0, 0; i < m && j < n; {
		freq := int(math.Min(float64(encoded1[i][1]), float64(encoded2[j][1])))
		val := encoded1[i][0] * encoded2[j][0]

		if len(res) == 0 || res[len(res)-1][0] != val {
			res = append(res, []int{val, freq})
		} else {
			res[len(res)-1][1] += freq
		}

		encoded1[i][1] -= freq
		encoded2[j][1] -= freq
		if encoded1[i][1] == 0 {
			i++
		}
		if encoded2[j][1] == 0 {
			j++
		}
	}
	return res
}
