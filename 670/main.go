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

// Method 2
func maximumSwap2(num int) int {
	if num == 0 {
		return 0
	}
	str := strconv.Itoa(num)
	data := []byte(str)

	max := 9
	for i := 0; i < len(data)-1; i++ {
		cur := int(data[i] - '0')
		for j := max; j > cur; j-- {
			for k := len(data) - 1; k > i; k-- {
				if int(data[k]-'0') == j {
					data[i], data[k] = data[k], data[i]
					val, _ := strconv.Atoi(string(data))
					return val
				}
			}

		}
		max = cur
	}
	return num
}
