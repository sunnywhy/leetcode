package main

// Largest Rectangle in Histogram
type stack []int

func (s stack) isEmpty() bool {
	return len(s) == 0
}
func (s stack) pop() (stack, int) {
	l := len(s)
	val := s[l-1]
	s = s[:l-1]
	return s, val
}
func (s stack) peek() int {
	l := len(s)
	return s[l-1]
}
func (s stack) push(val int) stack {
	s = append(s, val)
	return s
}
func largestRectangleArea(heights []int) int {
	s := stack{}
	var max, top int
	heights = append(heights, -1)
	for i, v := range heights {
		for !s.isEmpty() && v < heights[s.peek()] {
			s, top = s.pop()
			distance := 0
			if s.isEmpty() {
				distance = i
			} else {
				distance = i - s.peek() - 1
			}

			if max < distance*heights[top] {
				max = distance * heights[top]
			}
		}
		s = s.push(i)
	}

	return max
}
