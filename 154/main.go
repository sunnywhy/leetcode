package main

// Find Minimum in Rotated Sorted Array II
func findMin(nums []int) int {
	n := len(nums)
	left, right := 0, n-1
	for left <= right {
		mid := left + (right-left)/2
		if mid > 0 && nums[mid] < nums[mid-1] {
			return nums[mid]
		}
		//
		if nums[mid] < nums[right] { // right is sorted
			right = mid - 1
		} else if nums[mid] > nums[right] { //left is sorted
			left = mid + 1
		} else { // nums[mid] == nums[right]
			right--
		}
	}
	return nums[0]
}
