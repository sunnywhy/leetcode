package main

// Palindromic Substrings

// O(N^2)
func countSubstrings(s string) int {
	n := len(s)
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
	}

	var res int
	// loop the length of substring, from 1 to n
	for l := 1; l <= n; l++ {
		for i := 0; i <= n-l; i++ {
			j := i + l - 1
			if l == 1 {
				dp[i][j] = true
				res++
			} else if l == 2 {
				if s[i] == s[j] {
					dp[i][j] = true
					res++
				}
			} else {
				if s[i] == s[j] && dp[i+1][j-1] {
					dp[i][j] = true
					res++
				}
			}
		}
	}
	return res
}
