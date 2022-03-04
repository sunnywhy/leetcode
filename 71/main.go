package main

import "strings"

// Simplify Path

// O(N)
func simplifyPath(path string) string {
	arr := strings.Split(path, "/")
	var stack []string
	for _, v := range arr {
		if len(v) == 0 || v == "." {
			continue
		}
		if v == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, v)
		}
	}
	return "/" + strings.Join(stack, "/")
}
