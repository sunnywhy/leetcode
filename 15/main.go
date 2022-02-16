package main

import "sort"

// 3Sum
// use the Map way
func threeSum1(nums []int) [][]int {
	var res [][]int
	sort.Ints(nums)
	n := len(nums)
	data := make(map[int]int)

	for i, v := range nums {
		data[v] = i
	}

	for i := 0; i < n; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < n; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			target := 0 - nums[i] - nums[j]
			if k, ok := data[target]; ok {
				if k > j {
					res = append(res, []int{nums[i], nums[j], nums[k]})
				}
			}
		}
	}

	return res
}

// use the two pointers
func threeSum(nums []int) [][]int {
	var res [][]int
	sort.Ints(nums)
	n := len(nums)

	for i := 0; i < n; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j, k := i+1, n-1
		for j < k {
			if nums[i]+nums[j]+nums[k] == 0 {
				res = append(res, []int{nums[i], nums[j], nums[k]})
				for j+1 < n && nums[j+1] == nums[j] {
					j++
				}
				for k-1 >= 0 && nums[k-1] == nums[k] {
					k--
				}
				j++
				k--
			} else if nums[i]+nums[j]+nums[k] < 0 {
				j++
			} else {
				k--
			}
		}
	}
	return res
}
