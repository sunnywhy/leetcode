package main

// Single Number
func singleNumber(nums []int) int {
	data := make(map[int]bool)
	for _, v := range nums {
		exist := data[v]
		if exist {
			delete(data, v)
		} else {
			data[v] = true
		}
	}

	for k := range data {
		return k
	}
	return -1
}
