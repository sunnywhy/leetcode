package main

import "sort"

// Relative Sort Array
func relativeSortArray(arr1 []int, arr2 []int) []int {
	m2 := make(map[int]int)
	for idx, val := range arr2 {
		m2[val] = idx
	}

	sort.Slice(arr1, func(i, j int) bool {
		a, b := arr1[i], arr1[j]
		if idxA, ok := m2[a]; ok {
			if idxB, ok := m2[b]; ok {
				return idxA < idxB
			} else {
				return true
			}
		} else {
			if _, ok := m2[b]; ok {
				return false
			} else {
				return a < b
			}
		}
	})

	return arr1
}
