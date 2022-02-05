package main

// Valid Anagram
func isAnagram(s string, t string) bool {
	return count(s) == count(t)
}

func count(s string) [26]int {
	res := [26]int{}
	for _, v := range s {
		res[int(v-'a')]++
	}
	return res
}
