package main

import (
	"fmt"
)

// Kth Largest Element in an Array
func findKthLargest(nums []int, k int) int {
	quickSelect(nums, k, 0, len(nums)-1)
	return nums[k-1]
}
func quickSelect(nums []int, k, l, r int) {
	if l >= r {
		return
	}
	pos := partition(nums, l, r)
	if pos == k-1 {
		return
	} else if pos > k-1 {
		quickSelect(nums, k, l, pos-1)
	} else {
		quickSelect(nums, k, pos+1, r)
	}
}
func partition(nums []int, l, r int) int {
	pos, val := l, nums[r]
	for i := l; i < r; i++ {
		if nums[i] > val {
			nums[pos], nums[i] = nums[i], nums[pos]
			pos++
		}
	}
	nums[pos], nums[r] = nums[r], nums[pos]
	return pos
}

func main() {
	nums := []int{3, 2, 1, 5, 6, 4}
	fmt.Printf("expected = 5, actual = %d", findKthLargest(nums, 2))
}
