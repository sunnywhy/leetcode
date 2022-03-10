package main

import "strconv"

// Maximum Swap
func maximumSwap(num int) int {
	chars := []rune(strconv.Itoa(num))
	sourceIdx, destIdx := -1, -1
	rightMax, rightMaxIdx := int(chars[len(chars)-1]-'0'), len(chars)-1
	for i := len(chars) - 2; i >= 0; i-- {
		cur := int(chars[i] - '0')
		if cur < rightMax {
			sourceIdx = i
			destIdx = rightMaxIdx
		} else if cur > rightMax {
			rightMax = cur
			rightMaxIdx = i
		}
	}
	if sourceIdx != -1 {
		chars[sourceIdx], chars[destIdx] = chars[destIdx], chars[sourceIdx]
	}
	val, _ := strconv.Atoi(string(chars))
	return val
}
