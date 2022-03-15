package main

// Merge Sorted Array

func merge(num1 []int, m int, num2 []int, n int) {
	i, j, k := m-1, n-1, m+n-1

	for i >= 0 && j >= 0 {
		if num1[i] > num2[j] {
			num1[k] = num1[i]
			i--
		} else {
			num1[k] = num2[j]
			j--
		}
		k--
	}

	for j >= 0 {
		num1[k] = num2[j]
		j--
		k--
	}
}
