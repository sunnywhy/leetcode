package main

// Check if an Original String Exists Given Two Encoded Strings

// O(K*M*N), M, N is the length of s1 and s2, K is the possible biggest number in both strings
func possiblyEquals(s1 string, s2 string) bool {
	chars1, chars2 := []rune(s1), []rune(s2)
	memo := make(map[[3]int]bool) //i, j and diff
	return backtracking(chars1, chars2, 0, 0, 0, memo)
}

func backtracking(chars1, chars2 []rune, i, j, diff int, memo map[[3]int]bool) bool {
	m, n := len(chars1), len(chars2)
	if i == m && j == n {
		return diff == 0
	}

	if val, ok := memo[[3]int{i, j, diff}]; ok {
		return val
	}

	if i < m && j < n && diff == 0 && chars1[i] == chars2[j] && backtracking(chars1, chars2, i+1, j+1, 0, memo) {
		memo[[3]int{i, j, diff}] = true
		return true
	}

	if i < m && !isDigit(chars1[i]) && diff > 0 && backtracking(chars1, chars2, i+1, j, diff-1, memo) {
		memo[[3]int{i, j, diff}] = true
		return true
	}

	if j < n && !isDigit(chars2[j]) && diff < 0 && backtracking(chars1, chars2, i, j+1, diff+1, memo) {
		memo[[3]int{i, j, diff}] = true
		return true
	}

	var num int
	for k := i; k < m && isDigit(chars1[k]); k++ {
		num = num*10 + int(chars1[k]-'0')
		if backtracking(chars1, chars2, k+1, j, diff-num, memo) {
			memo[[3]int{i, j, diff}] = true
			return true
		}
	}

	num = 0
	for k := j; k < n && isDigit(chars2[k]); k++ {
		num = num*10 + int(chars2[k]-'0')
		if backtracking(chars1, chars2, i, k+1, diff+num, memo) {
			memo[[3]int{i, j, diff}] = true
			return true
		}
	}

	memo[[3]int{i, j, diff}] = false
	return false
}

func isDigit(c rune) bool {
	return c >= '0' && c <= '9'
}
