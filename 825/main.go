package main

// Friends Of Appropriate Ages

// O(N)
func numFriendRequests(ages []int) int {
	var buckets [121]int
	for _, v := range ages {
		buckets[v]++
	}

	var res int
	for i := 1; i < len(buckets); i++ {
		for j := 1; j <= i; j++ {
			if valid(i, j) {
				if i == j {
					res += buckets[i] * (buckets[j] - 1)
				} else {
					res += buckets[i] * buckets[j]
				}
			}
		}
	}
	return res
}

func valid(x, y int) bool {
	return !((y <= x/2+7) || y > x || (y > 100 && x < 100))
}
