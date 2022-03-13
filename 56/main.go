package main

import "sort"

// Merge Intervals
func merge(intervals [][]int) [][]int {
	var res [][]int

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	cur := intervals[0]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] > cur[1] {
			res = append(res, append([]int{}, cur...))
			cur = intervals[i]
		} else {
			if intervals[i][1] > cur[1] {
				cur[1] = intervals[i][1]
			}
		}
	}
	res = append(res, cur)
	return res
}
