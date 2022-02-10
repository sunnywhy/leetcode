package main

// Find First and Last Position of Element in Sorted Array
func searchRange(nums []int, target int) []int {
	res := make([]int, 2)
	res[0], res[1] = findFirst(nums, target), findLast(nums, target)
	return res
}
func findFirst(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			if mid == 0 || nums[mid-1] != target {
				return mid
			} else {
				right = mid - 1
			}
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}
func findLast(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			if mid == len(nums)-1 || nums[mid+1] != target {
				return mid
			} else {
				left = mid + 1
			}
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}
