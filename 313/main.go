package main

import "container/heap"

//  Super Ugly Number

type minHeap []int

func (h minHeap) Len() int {
	return len(h)
}
func (h minHeap) Less(i, j int) bool {
	return h[i] < h[j]
}
func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *minHeap) Pop() interface{} {
	n := len(*h)
	top := (*h)[n-1]
	*h = (*h)[:n-1]
	return top
}

func nthSuperUglyNumber(n int, primes []int) int {
	if n == 1 {
		return 1
	}
	h := &minHeap{}
	visited := make(map[int]bool)
	for _, v := range primes {
		heap.Push(h, v)
		visited[v] = true
	}
	for i := 2; i <= n; i++ {
		cur := heap.Pop(h).(int)
		if i == n {
			return cur
		}
		for _, v := range primes {
			if !visited[cur*v] {
				heap.Push(h, cur*v)
				visited[cur*v] = true
				break
			}
		}
	}
	return -1
}

func main() {
	primes := []int{2, 7, 13, 19}
	println(nthSuperUglyNumber(12, primes))
}
