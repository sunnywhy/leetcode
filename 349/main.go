package main

// Intersection of Two Arrays
func intersection(nums1 []int, nums2 []int) []int {
	m1 := make(map[int]bool)

	for _, v := range nums1 {
		m1[v] = true
	}

	temp := make(map[int]bool)
	for _, v := range nums2 {
		if m1[v] {
			temp[v] = true
		}
	}

	var res []int
	for k := range temp {
		res = append(res, k)
	}
	return res
}
