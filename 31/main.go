package main

// Next Permutation

// O(N)
func nextPermutation(nums []int) {
	pos := -1
	n := len(nums)
	for i := n - 1; i > 0; i-- {
		if nums[i-1] < nums[i] {
			pos = i - 1
			break
		}
	}

	if pos >= 0 {
		for i := n - 1; i >= 0; i-- {
			if nums[i] > nums[pos] {
				nums[i], nums[pos] = nums[pos], nums[i]
				break
			}
		}
	}

	for i, j := pos+1, n-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}
