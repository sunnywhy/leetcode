package main

// Search Insert Position
func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			if mid == len(nums)-1 || nums[mid+1] > target {
				return mid + 1
			} else {
				left = mid + 1
			}
		} else {
			if mid == 0 {
				return 0
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}
