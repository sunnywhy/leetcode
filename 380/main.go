package main

import "math/rand"

// Insert Delete GetRandom O(1)
// https://leetcode.com/problems/insert-delete-getrandom-o1/

type RandomizedSet struct {
	dataMap  map[int]int
	dataList []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		make(map[int]int),
		make([]int, 0),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.dataMap[val]; ok {
		return false
	}
	n := len(this.dataList)
	this.dataMap[val] = n
	this.dataList = append(this.dataList, val)
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	if idx, ok := this.dataMap[val]; ok {
		n := len(this.dataList)
		last := this.dataList[n-1]
		this.dataMap[last] = idx
		this.dataList[idx] = last
		this.dataList = this.dataList[:n-1]
		delete(this.dataMap, val)
		return true
	}
	return false
}

func (this *RandomizedSet) GetRandom() int {
	n := len(this.dataList)
	return this.dataList[rand.Intn(n)]
}
