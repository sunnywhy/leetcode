package main

// O(N*2^N
func subsets(nums []int) [][]int {
	n := len(nums)
	N := 1 << n
	res := make([][]int, N)
	for i := 0; i < N; i++ {
		cur := []int{}
		for j := 0; j < n; j++ {
			if i>>j&1 == 1 {
				cur = append(cur, nums[j])
			}
		}
		res[i] = cur
	}
	return res
}

// backtracking
// N * 2^N, copy array, the time complexity is O(N)
var res [][]int

func subsets2(nums []int) [][]int {
	res = nil
	cur := []int{}
	backtracking(nums, 0, &cur)
	return res
}
func backtracking(nums []int, pos int, cur *[]int) {
	if pos == len(nums) {
		tmp := make([]int, len(*cur))
		copy(tmp, *cur)
		res = append(res, tmp)
		return
	}

	backtracking(nums, pos+1, cur)
	*cur = append(*cur, nums[pos])
	backtracking(nums, pos+1, cur)
	*cur = (*cur)[:len(*cur)-1]

}
