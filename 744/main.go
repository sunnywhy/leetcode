package main

import "fmt"

// Find Smallest Letter Greater Than Target
func nextGreatestLetter(letters []byte, target byte) byte {
	left, right := 0, len(letters)-1
	for left <= right {
		mid := left + (right-left)/2
		if letters[mid] > target {
			if mid == 0 || letters[mid-1] <= target {
				return letters[mid]
			} else {
				right = mid - 1
			}
		} else {
			left = mid + 1
		}
	}
	return letters[0]
}

func main() {
	fmt.Printf("expected = c, output = %c \n", nextGreatestLetter([]byte{'c', 'f', 'j'}, 'a'))
	fmt.Printf("expected = a, output = %c", nextGreatestLetter([]byte{'a', 'b', 'j'}, 'm'))
}
