package main

// Remove Invalid Parentheses

var res []string

func removeInvalidParentheses(s string) []string {
	var left, right int
	for _, c := range s {
		if c == '(' {
			left++
		} else if c == ')' {
			if left > 0 {
				left--
			} else {
				right++
			}
		}
	}

	res = nil
	backtracking(s, left, right, 0)
	return res
}

func backtracking(s string, left, right int, pos int) {
	if left == 0 && right == 0 {
		if valid(s) {
			res = append(res, s)
		}
		return
	}
	for i := pos; i < len(s); i++ {
		if i > pos && s[i-1] == s[i] {
			continue
		}
		if right > 0 {
			if s[i] == ')' {
				backtracking(s[:i]+s[i+1:], left, right-1, i)
			}
		} else if left > 0 {
			if s[i] == '(' {
				backtracking(s[:i]+s[i+1:], left-1, right, i)
			}
		}
	}
}
func valid(s string) bool {
	var count int
	for _, c := range s {
		if c == '(' {
			count++
		} else if c == ')' {
			if count > 0 {
				count--
			} else {
				return false
			}
		}
	}
	return count == 0
}
