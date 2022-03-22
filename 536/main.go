package main

import "strconv"

// Construct Binary Tree from String

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func str2tree(s string) *TreeNode {
	if s == "" {
		return nil
	}
	n := len(s)
	leftStart, leftEnd := -1, -1
	var count int
	for i := 0; i < n; i++ {
		if s[i] == '(' {
			count++
			if leftStart == -1 {
				leftStart = i
			}
		} else if s[i] == ')' {
			count--
			if count == 0 {
				leftEnd = i
				break
			}
		}
	}

	if leftStart == -1 {
		val, _ := strconv.Atoi(s)
		return &TreeNode{Val: val}
	}

	val, _ := strconv.Atoi(s[:leftStart])
	node := &TreeNode{Val: val}
	node.Left = str2tree(s[leftStart+1 : leftEnd])
	if leftEnd < n-1 {
		node.Right = str2tree(s[leftEnd+2 : n-1])
	}
	return node
}
