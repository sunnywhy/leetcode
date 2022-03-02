package main

// Valid Palindrome II

func validPalindrome(s string) bool {
	if isValid, left, right := valid(s, 0, len(s)-1); isValid {
		return true
	} else {
		isLeftValid, _, _ := valid(s, left, right-1)
		isRightValid, _, _ := valid(s, left+1, right)
		return isLeftValid || isRightValid
	}
}

func valid(s string, start, end int) (bool, int, int) {
	for start < end {
		if s[start] != s[end] {
			return false, start, end
		}
		start++
		end--
	}

	return true, start, end
}
