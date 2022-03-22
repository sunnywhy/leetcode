package main

// Max Consecutive Ones III

func longestOnes(nums []int, k int) int {
	n := len(nums)
	var left, right, res int
	for right < n {
		if nums[right] == 0 {
			k--
		}
		right++
		for k < 0 {
			if nums[left] == 0 {
				k++
			}
			left++
		}
		if res < right-left {
			res = right - left
		}
	}
	return res
}
