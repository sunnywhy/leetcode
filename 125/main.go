package main

import "unicode"

// Valid Palindrome

func isPalindrome(s string) bool {
	chars := []rune(s)

	i, j := 0, len(chars)-1
	for i < j {
		if !isAlphanumber(chars[i]) {
			i++
			continue
		}
		if !isAlphanumber(chars[j]) {
			j--
			continue
		}
		if unicode.ToUpper(chars[i]) != unicode.ToUpper(chars[j]) {
			return false
		}
		i++
		j--
	}
	return true
}
func isAlphanumber(c rune) bool {
	if c >= '0' && c <= '9' {
		return true
	}
	if c >= 'a' && c <= 'z' {
		return true
	}
	if c >= 'A' && c <= 'Z' {
		return true
	}
	return false
}
