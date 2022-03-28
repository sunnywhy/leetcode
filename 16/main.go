package main

import (
	"math"
	"sort"
)

// 3Sum Closest

func threeSumClosest(nums []int, target int) int {
	diff := math.MaxFloat64
	var result int
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		for j, k := i+1, len(nums)-1; j < k; {
			sum := nums[i] + nums[j] + nums[k]
			if diff > math.Abs(float64(sum-target)) {
				diff = math.Abs(float64(sum - target))
				result = sum
			}
			if sum < target {
				j++
			} else {
				k--
			}
		}
	}
	return result
}
