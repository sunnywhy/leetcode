package main

import "fmt"

// Search in Rotated Sorted Array
func search(nums []int, target int) int {
	l, r := 0, len(nums)-1

	for l <= r {
		m := l + (r-l)/2
		if nums[m] == target {
			return m
		} else if nums[l] <= nums[m] { // left is sorted
			if target >= nums[l] && target < nums[m] {
				r = m - 1
			} else {
				l = m + 1
			}
		} else { // right is sorted
			if target > nums[m] && target <= nums[r] {
				l = m + 1
			} else {
				r = m - 1
			}
		}
	}
	return -1
}

func main() {
	fmt.Printf("expected = 4, actual = %d", search([]int{4, 5, 6, 7, 0, 1, 2}, 0))
}
