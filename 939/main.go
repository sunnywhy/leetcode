package main

import "math"

// Minimum Area Rectangle
// https://leetcode.com/problems/minimum-area-rectangle/

func minAreaRect(points [][]int) int {
	n := len(points)
	data := make(map[[2]int]bool)

	for _, p := range points {
		data[[2]int{p[0], p[1]}] = true
	}

	res := math.MaxInt
	for i := 0; i < n-1; i++ {
		x1, y1 := points[i][0], points[i][1]
		for j := i + 1; j < n; j++ {
			x2, y2 := points[j][0], points[j][1]
			if data[[2]int{x1, y2}] && data[[2]int{x2, y1}] {
				area := (x1 - x2) * (y1 - y2)
				if area < 0 {
					area = -area
				}
				if area != 0 && res > area {
					res = area
				}
			}
		}
	}
	if res == math.MaxInt {
		return 0
	}
	return res
}
