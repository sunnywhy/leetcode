package main

import "strings"

// Word Break II
func wordBreak(s string, wordDict []string) []string {
	wordMap := make(map[string]bool)
	for _, w := range wordDict {
		wordMap[w] = true
	}

	var res, cur []string
	backtracking([]rune(s), wordMap, 0, &cur, &res)
	return res
}

func backtracking(chars []rune, wordMap map[string]bool, start int, cur, res *[]string) {
	if start == len(chars) {
		*res = append(*res, strings.Join(*cur, " "))
		return
	}

	for i := start; i < len(chars); i++ {
		w := string(chars[start : i+1])
		if !wordMap[w] {
			continue
		}

		*cur = append(*cur, w)
		backtracking(chars, wordMap, i+1, cur, res)
		*cur = (*cur)[:len(*cur)-1]
	}
}
