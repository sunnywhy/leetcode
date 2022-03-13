package main

// Continuous Subarray Sum

func checkSubarraySum(nums []int, k int) bool {
	m := make(map[int]int)
	m[0] = -1
	var total int

	for i, v := range nums {
		total += v
		if pos, ok := m[total%k]; ok {
			if i-pos > 1 {
				return true
			}
		} else {
			m[total%k] = i
		}
	}
	return false
}
