package main

// Two Sum
func twoSum(nums []int, target int) []int {
	data := make(map[int]int)
	for idx, val := range nums {
		if i, ok := data[target-val]; ok {
			return []int{i, idx}
		} else {
			data[val] = idx
		}
	}
	return []int{-1, -1}
}
