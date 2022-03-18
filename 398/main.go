package main

import "math/rand"

// Random Pick Index

type Solution struct {
	data []int
}

func Constructor(nums []int) Solution {
	return Solution{nums}
}

func (this *Solution) Pick(target int) int {
	var count int
	var res int
	for i, v := range this.data {
		if v == target {
			count++
			if rand.Intn(count) == 0 {
				res = i
			}
		}
	}
	return res
}

// Method 2: Map
type Solution2 struct {
	data map[int][]int
}

func Constructor2(nums []int) Solution2 {
	data := make(map[int][]int)
	for i, v := range nums {
		data[v] = append(data[v], i)
	}
	return Solution2{data}
}

func (this *Solution2) Pick(target int) int {
	idxList := this.data[target]
	r := rand.Intn(len(idxList))
	return idxList[r]
}
