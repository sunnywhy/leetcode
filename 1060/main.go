package main

// Missing Element in Sorted Array
// https://leetcode.com/problems/missing-element-in-sorted-array/

func missingElement(nums []int, k int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		m := l + (r-l)/2
		if nums[m] < nums[0]+m+k {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return nums[0] + l + k - 1
}
