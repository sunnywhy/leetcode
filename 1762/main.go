package main

// Buildings With an Ocean View

// O(N)
func findBuildings(heights []int) []int {
	var res []int
	var max int
	for i := len(heights) - 1; i >= 0; i-- {
		if heights[i] > max {
			res = append(res, i)
			max = heights[i]
		}
	}

	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}
