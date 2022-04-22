package main

// Find First and Last Position of Element in Sorted Array
func searchRange(nums []int, target int) []int {
	first := findFirst(nums, target)
	last := findLast(nums, target)
	return []int{first, last}
}

func findFirst(nums []int, target int) int {
	l, r := 0, len(nums)
	for l < r {
		m := l + (r-l)/2
		if nums[m] >= target {
			r = m
		} else {
			l = m + 1
		}
	}
	if l < len(nums) && nums[l] == target {
		return l
	}
	return -1
}
func findLast(nums []int, target int) int {
	l, r := 0, len(nums)
	for l < r {
		m := l + (r-l)/2
		if nums[m] > target {
			r = m
		} else {
			l = m + 1
		}
	}
	if l > 0 && nums[l-1] == target {
		return l - 1
	}
	return -1
}
