package main

// Remove Invalid Parentheses

var res []string

func removeInvalidParentheses(s string) []string {
	res = nil
	var lefts, rights int
	for _, c := range s {
		if c == '(' {
			lefts++
		} else if c == ')' {
			if lefts > 0 {
				lefts--
			} else {
				rights++
			}
		}
	}

	backtracking(s, 0, lefts, rights)

	return res
}

func backtracking(s string, start, lefts, rights int) {
	if lefts == 0 && rights == 0 {
		if valid(s) {
			res = append(res, s)
		}
		return
	}

	for i := start; i < len(s); i++ {
		if i > start && s[i] == s[i-1] {
			continue
		}
		if s[i] == ')' && rights > 0 {
			chars := []rune(s)
			cur := string(append(chars[:i], chars[i+1:]...))
			backtracking(cur, i, lefts, rights-1)
		} else if s[i] == '(' && lefts > 0 {
			chars := []rune(s)
			cur := string(append(chars[:i], chars[i+1:]...))
			backtracking(cur, i, lefts-1, rights)
		}
	}
}

func valid(s string) bool {
	var count int
	for _, c := range s {
		if c == '(' {
			count++
		} else if c == ')' {
			if count == 0 {
				return false
			}
			count--
		}
	}
	return count == 0
}
