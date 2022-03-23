package main

// Maximum Size Subarray Sum Equals k

func maxSubArrayLen(nums []int, k int) int {
	data := make(map[int]int) // key is preSum, value is index
	data[0] = -1
	var total, res int
	for i, v := range nums {
		total += v
		if startPos, ok := data[total-k]; ok {
			if res < i-startPos {
				res = i - startPos
			}
		}
		if _, ok := data[total]; !ok {
			data[total] = i
		}
	}
	return res
}
