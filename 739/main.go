package main

// Daily temperatures

func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	stack := []int{}
	answers := make([]int, n)
	for k, v := range temperatures {
		if len(stack) == 0 || temperatures[stack[len(stack)-1]] >= v {
			stack = append(stack, k)
			continue
		}
		for len(stack) > 0 && temperatures[stack[len(stack)-1]] < v {
			idx := stack[len(stack)-1]
			answers[idx] = k - idx
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, k)
	}
	return answers
}
