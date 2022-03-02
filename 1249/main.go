package main

// Minimum Remove to Make Valid Parentheses

// O(N) : scan twice, without change the original s
func minRemoveToMakeValid(s string) string {
	var lefts []int
	var rights []int
	for idx, ch := range s {
		if ch == '(' {
			lefts = append(lefts, idx)
		} else if ch == ')' {
			if len(lefts) > 0 {
				lefts = lefts[:len(lefts)-1]
			} else {
				rights = append(rights, idx)
			}
		}
	}
	skips := make(map[int]bool)
	for _, i := range lefts {
		skips[i] = true
	}
	for _, i := range rights {
		skips[i] = true
	}
	var res []rune
	for idx, ch := range s {
		if skips[idx] {
			continue
		}
		res = append(res, ch)
	}
	return string(res)
}

// O(N) also scan twice, change the original s
func minRemoveToMakeValid2(s string) string {
	source := []rune(s)

	var count int
	for i, c := range source {
		if c == '(' {
			count++
		} else if c == ')' {
			if count > 0 {
				count--
			} else {
				source[i] = '*'
			}
		}
	}
	count = 0
	for i := len(source) - 1; i >= 0; i-- {
		if source[i] == ')' {
			count++
		} else if source[i] == '(' {
			if count > 0 {
				count--
			} else {
				source[i] = '*'
			}
		}
	}
	var res []rune
	for _, ch := range source {
		if ch == '*' {
			continue
		}
		res = append(res, ch)
	}
	return string(res)
}
