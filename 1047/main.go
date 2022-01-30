package main

import "strings"

// Remove All Adjacent Duplicates In String
func removeDuplicates(s string) string {
	stack := []rune{}
	for _, v := range s {
		l := len(stack)
		if l == 0 || v != stack[l-1] {
			stack = append(stack, v)
		}else {
			stack = stack[:l-1]
		}
	}
	builder := strings.Builder{}
	for _, v := range stack {
		builder.WriteRune(v)
	}
	return builder.String()
}