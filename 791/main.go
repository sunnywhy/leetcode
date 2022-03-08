package main

import "sort"

// Custom Sort String

// Method 1: custom sorting, O(NlogN)
func customSortString(order string, s string) string {
	orderMap := make(map[rune]int)
	for i, c := range order {
		orderMap[c] = i
	}
	chars := []rune(s)
	sort.Slice(chars, func(i, j int) bool {
		return orderMap[chars[i]] < orderMap[chars[j]]
	})
	return string(chars)
}

// O(N)
func customSortString2(order string, s string) string {
	data := make([]int, 26)
	for _, c := range s {
		data[c-'a']++
	}
	var res []rune
	for _, o := range order {
		for i := 0; i < data[o-'a']; i++ {
			res = append(res, o)
		}
		data[o-'a'] = 0
	}
	for i, c := range data {
		for j := 0; j < c; j++ {
			res = append(res, rune('a'+i))
		}
	}
	return string(res)
}
