package main

import "math/rand"

// Random Pick with Weight

type Solution struct {
	Sums  []int
	Total int
}

func Constructor(w []int) Solution {
	sums := make([]int, len(w))
	var total int
	for i, v := range w {
		total += v
		sums[i] = total
	}
	return Solution{sums, total}
}

func (this *Solution) PickIndex() int {
	random := rand.Intn(this.Total) // return [0, total)
	left, right := 0, len(this.Sums)
	for left <= right {
		mid := left + (right-left)/2
		if this.Sums[mid] > random {
			if mid == 0 || this.Sums[mid-1] <= random {
				return mid
			} else {
				right = mid - 1
			}
		} else {
			left = mid + 1
		}
	}
	return -1
}
