package main

import "math"

// String to Integer (atoi)
// https://leetcode.com/problems/string-to-integer-atoi/

func myAtoi(s string) int {
	sign := 1 // 1: positive, -1: negative
	i := 0
	for i < len(s) && s[i] == ' ' {
		i++
	}
	if i == len(s) {
		return 0
	}
	if s[i] == '-' {
		sign = -1
	}
	if s[i] == '-' || s[i] == '+' {
		i++
	}
	var res int
	for i < len(s) && s[i] >= '0' && s[i] <= '9' {
		num := int(s[i] - '0')
		if (math.MaxInt32-num)/10 < res {
			if sign == 1 {
				return math.MaxInt32
			} else {
				return math.MinInt32
			}
		}
		res = res*10 + num
		i++
	}
	return res * sign
}
