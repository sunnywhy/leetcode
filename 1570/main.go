package main

// Dot Product of Two Sparse Vectors

type SparseVector struct {
	Data map[int]int
}

func Constructor(nums []int) SparseVector {
	data := make(map[int]int)
	for i, v := range nums {
		data[i] = v
	}
	return SparseVector{data}
}

func (this *SparseVector) dotProduct(vec SparseVector) int {
	l1, l2 := len(this.Data), len(vec.Data)
	if l1 > l2 {
		return vec.dotProduct(*this)
	}
	var res int
	for k, v := range this.Data {
		res += v * vec.Data[k]
	}
	return res
}
