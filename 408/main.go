package main

import "strconv"

// Valid Word Abbreviation

func validWordAbbreviation(word string, abbr string) bool {
	l1, l2 := len(word), len(abbr)
	i, j := 0, 0
	for i < l1 && j < l2 {
		if isDigit(abbr[j]) {
			startPos := j
			for j < l2 && isDigit(abbr[j]) {
				j++
			}
			if abbr[startPos] == '0' {
				return false
			}
			count, _ := strconv.Atoi(string(abbr[startPos:j]))
			i += count
		} else {
			if word[i] != abbr[j] {
				return false
			}
			i++
			j++
		}
	}
	return i == l1 && j == l2
}

func isDigit(ch byte) bool {
	if ch >= '0' && ch <= '9' {
		return true
	}
	return false
}
