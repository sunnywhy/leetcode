package main

// Capacity to Ship Packages Within D Days

func shipWithinDays(weights []int, days int) int {
	var max, total int
	for _, v := range weights {
		if v > max {
			max = v
		}
		total += v
	}

	l, r := max, total
	for l <= r {
		m := l + (r-l)/2
		needed := calculate(weights, m)
		if needed <= days {
			if m == max || calculate(weights, m-1) > days {
				return m
			} else {
				r = m - 1
			}
		} else {
			l = m + 1
		}
	}
	return -1
}

func calculate(weights []int, capacity int) int {
	var days, total int

	for _, v := range weights {
		if total+v > capacity {
			days++
			total = v
		} else {
			total += v
		}
	}
	if total > 0 {
		days++
	}
	return days
}
