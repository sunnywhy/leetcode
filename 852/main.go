package main

// Peak Index in a Mountain Array
func peakIndexInMountainArray(arr []int) int {
	left, right := 0, len(arr)-1
	for left <= right {
		mid := left + (right-left)/2
		if mid == 0 {
			left = mid + 1
		} else if mid == len(arr)-1 {
			right = mid - 1
		} else if arr[mid] > arr[mid-1] && arr[mid] > arr[mid+1] {
			return mid
		} else if arr[mid] > arr[mid-1] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}
