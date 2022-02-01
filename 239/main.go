package main

import "container/list"
// Sliding Window Maximum

func maxSlidingWindow(nums []int, k int) []int {
    res := []int{}
	q := list.New()
	for i, v := range nums {
		for q.Len() > 0 && v > nums[q.Back().Value.(int)] {
			q.Remove(q.Back())
		}
		q.PushBack(i)
		if q.Front().Value.(int) <= i - k {
			q.Remove(q.Front())
		}
		if i >= k - 1 {
			res = append(res, nums[q.Front().Value.(int)])
		}
	}
	return res
}