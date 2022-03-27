package main

import "math"

// Minimum Window Substring
// https://leetcode.com/problems/minimum-window-substring/

func minWindow(s string, t string) string {
	data := make([]int, 128)
	var left, right int
	count := len(t)
	minLen := math.MaxInt
	start := -1

	for i := 0; i < len(t); i++ {
		data[int(t[i])]++
	}

	for right < len(s) {
		c := s[right]
		if data[int(c)] > 0 {
			count--
		}
		data[int(c)]--

		for count == 0 {
			if minLen > right-left+1 {
				minLen = right - left + 1
				start = left
			}

			data[int(s[left])]++
			if data[int(s[left])] > 0 {
				count++
			}
			left++
		}
		right++
	}

	if start == -1 {
		return ""
	}
	return s[start : start+minLen]
}
