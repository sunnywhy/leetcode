package main

// Can Place Flowers
// https://leetcode.com/problems/can-place-flowers/

func canPlaceFlowers(flowerbed []int, n int) bool {
	var count int
	for i := 0; i < len(flowerbed); i++ {
		var flag1, flag2 bool
		if flowerbed[i] == 1 {
			continue
		}
		if i == 0 || flowerbed[i-1] == 0 {
			flag1 = true
		}
		if i == len(flowerbed)-1 || flowerbed[i+1] == 0 {
			flag2 = true
		}
		if flag1 && flag2 {
			flowerbed[i] = 1
			count++
		}
	}
	return count >= n
}
