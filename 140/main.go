package main

import "strings"

// Word Break II
var res []string

func wordBreak(s string, wordDict []string) []string {
	res = nil
	wordMap := make(map[string]bool)
	for _, word := range wordDict {
		wordMap[word] = true
	}
	var cur []string
	backtracking(s, wordMap, 0, &cur)
	return res
}
func backtracking(s string, wordMap map[string]bool, pos int, cur *[]string) {
	if pos >= len(s) {
		res = append(res, strings.Join(*cur, " "))
		return
	}
	for i := pos; i < len(s); i++ {
		temp := s[pos : i+1]
		if wordMap[temp] {
			*cur = append(*cur, temp)
			backtracking(s, wordMap, i+1, cur)
			*cur = (*cur)[:len(*cur)-1]
		}
	}
}
