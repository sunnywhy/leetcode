package main

import "strings"

// Multiply Strings
// https://leetcode.com/problems/multiply-strings/
func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	m, n := len(num1), len(num2)
	data := make([]int, m+n)

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			cur := data[i+j+1] + int(num1[i]-'0')*int(num2[j]-'0')
			data[i+j+1] = cur % 10
			data[i+j] += cur / 10
		}
	}

	var sb strings.Builder
	for i, v := range data {
		if i == 0 && v == 0 {
			continue
		}
		sb.WriteByte(byte(v + '0'))
	}
	return sb.String()
}
