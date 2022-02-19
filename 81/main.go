package main

// Search in Rotated Sorted Array II
func search(nums []int, target int) bool {
	n := len(nums)
	left, right := 0, n-1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			return true
		}
		// special case [2,1,2,2,2]
		if nums[mid] == nums[left] {
			left++
			continue
		}

		if nums[left] <= nums[mid] { // left is sorted
			if target >= nums[left] && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else { // right is sorted
			if target > nums[mid] && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return false
}
