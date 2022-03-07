package main

// Subarray Sum Equals K

//O(N)
func subarraySum(nums []int, k int) int {
	var res, total int
	m := make(map[int]int)
	m[0] = 1
	for _, v := range nums {
		total += v
		res += m[total-k]
		m[total]++
	}
	return res
}
