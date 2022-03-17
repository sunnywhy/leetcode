package main

// Kth Missing Positive Number

func findKthPositive(arr []int, k int) int {
	l, r := 0, len(arr)-1
	for l <= r {
		m := l + (r-l)/2
		if arr[m]-m-1 < k {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return l + k
}
