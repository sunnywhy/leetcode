package main

// Group Anagrams
func groupAnagrams(strs []string) [][]string {
	data := make(map[[26]int][]string)
	for _, s := range strs {
		counts := letterCounts(s)
		if val, ok := data[counts]; ok {
			data[counts] = append(val, s)
		} else {
			data[counts] = []string{s}
		}
	}
	var res [][]string
	for _, v := range data {
		res = append(res, v)
	}
	return res
}

func letterCounts(str string) [26]int {
	counts := [26]int{}
	for _, v := range []byte(str) {
		counts[int(v-'a')]++
	}
	return counts
}
