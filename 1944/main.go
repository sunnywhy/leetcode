package main

// Number of Visible People in a Queue
// https://leetcode.com/problems/number-of-visible-nodes-in-the-binary-tree/

func canSeePersonsCount(heights []int) []int {
	n := len(heights)
	res := make([]int, n)
	stack := make([]int, 0)

	for i, h := range heights {
		for len(stack) > 0 && heights[stack[len(stack)-1]] < h {
			res[stack[len(stack)-1]]++
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			res[stack[len(stack)-1]]++
		}
		stack = append(stack, i)
	}
	return res
}
