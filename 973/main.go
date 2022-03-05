package main

// K Closest Points to Origin

func kClosest(points [][]int, k int) [][]int {
	quickSelect(&points, k, 0, len(points)-1)
	return points[:k]
}
func quickSelect(points *[][]int, k, start, end int) {
	if start > end {
		return
	}
	pos := partition(points, start, end)
	if pos == k-1 {
		return
	} else if pos > k-1 {
		quickSelect(points, k, start, pos-1)
	} else {
		quickSelect(points, k, pos+1, end)
	}
}

func partition(points *[][]int, start, end int) int {
	pos := start
	val := (*points)[end]
	for i := start; i <= end; i++ {
		cur := (*points)[i]
		if cur[0]*cur[0]+cur[1]*cur[1] < val[0]*val[0]+val[1]*val[1] {
			(*points)[pos], (*points)[i] = (*points)[i], (*points)[pos]
			pos++
		}
	}
	(*points)[pos], (*points)[end] = (*points)[end], (*points)[pos]
	return pos
}
