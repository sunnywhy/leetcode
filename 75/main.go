package main

// Sort Colors
func sortColors(nums []int) {
	quickSort(nums, 0, len(nums)-1)
}
func quickSort(nums []int, left int, right int) {
	if left >= right {
		return
	}
	pos := partition(nums, left, right)
	quickSort(nums, left, pos-1)
	quickSort(nums, pos+1, right)
}
func partition(nums []int, left int, right int) int {
	pos, val := left, nums[right]
	for i := left; i < right; i++ {
		if nums[i] < val {
			nums[pos], nums[i] = nums[i], nums[pos]
			pos++
		}
	}
	nums[pos], nums[right] = nums[right], nums[pos]
	return pos
}
