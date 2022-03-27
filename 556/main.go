package main

import (
	"math"
	"strconv"
)

// Next Greater Element III
// https://leetcode.com/problems/next-greater-element-iii/

func nextGreaterElement(n int) int {
	b := []byte(strconv.Itoa(n))
	start := -1
	for i := len(b) - 1; i > 0; i-- {
		if b[i] > b[i-1] {
			start = i - 1
			break
		}
	}
	if start == -1 {
		return -1
	}
	for i := len(b) - 1; i > start; i-- {
		if b[i] > b[start] {
			b[start], b[i] = b[i], b[start]
			break
		}
	}
	for i, j := start+1, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	var res int
	for _, v := range b {
		num := int(v - '0')
		if (math.MaxInt32-num)/10 < res {
			return -1
		}
		res = res*10 + num
	}
	return res
}
