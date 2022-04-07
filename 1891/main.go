package main

// O(NLogN), easy to understand
func maxLength(ribbons []int, k int) int {
	var max int
	for _, v := range ribbons {
		if v > max {
			max = v
		}
	}
	l, r := 1, max

	for l <= r {
		m := l + (r-l)/2
		val := getRibbons(ribbons, m)
		if val >= k {
			if getRibbons(ribbons, m+1) < k {
				return m
			} else {
				l = m + 1
			}
		} else {
			r = m - 1
		}
	}
	return 0
}
func getRibbons(ribbons []int, length int) int {
	var total int
	for _, v := range ribbons {
		total += v / length
	}
	return total
}

// O(N*LogN), shorter
func maxLength2(ribbons []int, k int) int {
	var max int
	for _, v := range ribbons {
		if v > max {
			max = v
		}
	}

	l, r := 1, max+1
	for l < r {
		m := l + (r-l)/2
		if !valid(ribbons, k, m) {
			r = m
		} else {
			l = m + 1
		}
	}
	return l - 1
}
func valid(ribbons []int, k int, size int) bool {
	var total int
	for _, v := range ribbons {
		total += v / size
	}
	return total >= k
}
