package main

import "math"

// Interval List Intersections

func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	m, n := len(firstList), len(secondList)
	var i, j int
	var res [][]int
	for i < m && j < n {
		maxStart := int(math.Max(float64(firstList[i][0]), float64(secondList[j][0])))
		minEnd := int(math.Min(float64(firstList[i][1]), float64(secondList[j][1])))
		if maxStart <= minEnd {
			res = append(res, []int{maxStart, minEnd})
		}

		if firstList[i][1] < secondList[j][1] {
			i++
		} else {
			j++
		}
	}
	return res
}
