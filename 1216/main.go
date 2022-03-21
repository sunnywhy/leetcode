package main

import "math"

// Valid Palindrome III

func isValidPalindrome(s string, k int) bool {
	longest := longestPalindrome(&s)
	return longest >= len(s)-k
}

func longestPalindrome(s *string) int {
	n := len(*s)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}

	chars := []byte(*s)
	for l := 1; l <= n; l++ {
		for i := 0; i <= n-l; i++ {
			j := i + l - 1

			if l == 1 {
				dp[i][j] = 1
			} else {
				if chars[i] == chars[j] {
					dp[i][j] = dp[i+1][j-1] + 2
				} else {
					dp[i][j] = int(math.Max(float64(dp[i+1][j]), float64(dp[i][j-1])))
				}
			}
		}
	}
	return dp[0][n-1]
}
