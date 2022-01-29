package main

// Valid Parentheses
func isValid(s string) bool {
	stack := []rune{}
	for _, v := range s {
		if v == '(' || v == '{' || v == '[' {
			stack = append(stack, v)
		}else {
			l := len(stack)
			if l == 0 {
				return false
			}
			top := stack[l-1]
			if v == ')' && top != '(' {
				return false
			}
			if v == '}' && top != '{' {
				return false
			}
			if v == ']' && top != '[' {
				return false
			}
			stack = stack[:l-1]
		}
	}
	return len(stack) == 0
}

func main() {
	println(isValid("{}[[()]]"))
	println(isValid("({}[[()]]"))
}