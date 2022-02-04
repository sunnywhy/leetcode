package main

//  Trapping Rain Water

// method1: keep tracing the left Max and the right Max
func trap(height []int) int {
	n := len(height)
	lefts := make([]int, n)
	leftMax, rightMax := 0, 0
	for i, v := range height {
		if v > leftMax {
			leftMax = v
		}
		lefts[i] = leftMax
	}
	var result int
	for i := n - 1; i >= 0; i-- {
		if height[i] > rightMax {
			rightMax = height[i]
		}
		result += min(lefts[i], rightMax) - height[i]
	}
	return result
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

// use a stack
func trap1(height []int) int {
	stack := []int{}
	var res int
	for i, v := range height {
		for len(stack) > 0 && v > height[stack[len(stack)-1]] {
			top := height[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			leftBound := stack[len(stack)-1]

			res += (i - leftBound - 1) * (min(v, height[leftBound]) - top)
		}
		stack = append(stack, i)
	}
	return res
}

// two pointers
func trap2(height []int) int {
	n := len(height)
	leftMax, rightMax := height[0], height[n-1]
	var res int
	for i, j := 0, n-1; i < j; {
		if height[i] < height[j] {
			leftMax = max(leftMax, height[i])
			res += leftMax - height[i]
			i++
		} else {
			rightMax = max(rightMax, height[j])
			res += rightMax - height[j]
			j--
		}
	}
	return res
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
