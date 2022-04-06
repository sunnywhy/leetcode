package main

// Find Peak Element
// easy to understand
func findPeakElement(nums []int) int {
	n := len(nums)
	if n == 1 {
		return 0
	}
	left, right := 0, n-1
	for left <= right {
		mid := left + (right-left)/2
		if mid == 0 {
			if nums[mid] > nums[mid+1] {
				return mid
			} else {
				left = mid + 1
			}
		} else if mid == n-1 {
			if nums[mid] > nums[mid-1] {
				return mid
			} else {
				right = mid - 1
			}
		} else if nums[mid] > nums[mid-1] && nums[mid] > nums[mid+1] {
			return mid
		} else if nums[mid] < nums[mid-1] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

// shorter
func findPeakElement2(nums []int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		if l == r {
			return l
		}
		m := l + (r-l)/2
		if nums[m] < nums[m+1] {
			l = m + 1
		} else {
			r = m
		}
	}
	return -1
}
