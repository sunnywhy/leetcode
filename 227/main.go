package main

// Basic Calculator II
func calculate(s string) int { 
	stack := []int{}
	n := len(s)
	val := 0
	operator := '+'
	for i, v := range s {
		if(isDigit(v)) {
			val = val * 10 + int(v-'0')
			if i != n -1 {
				continue
			}
		}
		if v == ' ' && i != n - 1 {
			continue
		}
		switch operator {
		case '+':
			stack = append(stack, val)
		case '-':
			stack = append(stack, -val)
		case '*':
			stack[len(stack)-1] = stack[len(stack)-1] * val
		case '/':
			stack[len(stack)-1] = stack[len(stack)-1] / val
		}
		val = 0
		operator = v
	}
	result := 0
	for _, v := range stack {
		result += v
	}
	return result
}
func isDigit(c rune) bool {
	return c >= '0' && c <= '9'
}
