package main

import "math"

// Sum of k-Mirror Numbers
// https://leetcode.com/problems/sum-of-k-mirror-numbers/
func kMirror(k int, n int) int64 {
	var res int64
	var count int
	for l := 1; l < 100; l++ { // l is the length of the number
		for i := int64(math.Pow10(l - 1)); i < int64(math.Pow10(l)); i++ {
			num := generatePalindrome10(i, true)
			if isKMirror(num, k) {
				res += num
				count++
			}
			if count == n {
				return res
			}
		}
		for i := int64(math.Pow10(l - 1)); i < int64(math.Pow10(l)); i++ {
			num := generatePalindrome10(i, false)
			if isKMirror(num, k) {
				res += num
				count++
			}
			if count == n {
				return res
			}
		}

	}
	return 0
}
func generatePalindrome10(a int64, keepMiddle bool) int64 {
	b := a
	if keepMiddle {
		b /= 10
	}
	var reversed int64
	var count int
	for b > 0 {
		reversed = reversed*10 + b%10
		b /= 10
		count++
	}
	return a*int64(math.Pow10(count)) + reversed
}

var data [100]int64

func isKMirror(a int64, k int) bool {
	var count int
	for a > 0 {
		data[count] = a % int64(k)
		a /= int64(k)
		count++
	}
	for i, j := 0, count-1; i < j; i, j = i+1, j-1 {
		if data[i] != data[j] {
			return false
		}
	}
	return true
}
func main() {
	var k int
	var n int
	k, n = 2, 5
	println(kMirror(k, n))
}
