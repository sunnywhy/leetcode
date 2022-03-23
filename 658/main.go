package main

// Find K Closest Elements

func findClosestElements(arr []int, k int, x int) []int {
	i := closestSmallerOrEqual(arr, x)
	j := i + 1

	for k > 0 {
		if i < 0 {
			j++
		} else if j >= len(arr) {
			i--
		} else if x-arr[i] > arr[j]-x {
			j++
		} else {
			i--
		}
		k--
	}
	return arr[i+1 : j]
}

func closestSmallerOrEqual(arr []int, x int) int {
	l, r := 0, len(arr)-1
	for l <= r {
		m := l + (r-l)/2
		if arr[m] <= x {
			if m == len(arr)-1 || arr[m+1] > x {
				return m
			} else {
				l = m + 1
			}
		} else {
			r = m - 1
		}
	}
	return -1
}
