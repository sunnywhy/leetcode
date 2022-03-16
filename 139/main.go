package main

// Word Break

func wordBreak(s string, wordDict []string) bool {
	wordMap := make(map[string]bool)
	for _, w := range wordDict {
		wordMap[w] = true
	}

	chars := []rune(s)
	n := len(chars)
	dp := make([]bool, n+1)
	dp[0] = true

	for i := 1; i <= n; i++ { // i is the length of the string we validate
		for j := 0; j < i; j++ {
			if dp[j] && wordMap[string(chars[j:i])] {
				dp[i] = true
				break
			}
		}
	}

	return dp[n]
}
