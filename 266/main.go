package main

// Palindrome Permutation

func canPermutePalindrome(s string) bool {
	var data [26]int
	for _, c := range s {
		data[int(c-'a')]++
	}

	var count int
	for _, v := range data {
		if v%2 == 1 {
			count++
		}
	}
	return count <= 1
}
