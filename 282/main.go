package main

import "strconv"

// Expression Add Operators

var res []string

func addOperators(num string, target int) []string {
	res = nil
	backtracking(num, target, 0, 0, 0, "")
	return res
}

func backtracking(num string, target, start, cur, prev int, data string) {
	if start == len(num) {
		if cur == target {
			res = append(res, data)
		}
		return
	}

	for i := start; i < len(num); i++ {
		temp := num[start : i+1]
		if len(temp) > 1 && temp[0] == '0' {
			continue
		}
		val, _ := strconv.Atoi(temp)
		if start == 0 {
			backtracking(num, target, i+1, val, val, temp)
		} else {
			backtracking(num, target, i+1, cur+val, val, data+"+"+temp)
			backtracking(num, target, i+1, cur-val, -val, data+"-"+temp)
			backtracking(num, target, i+1, cur-prev+prev*val, prev*val, data+"*"+temp)
		}
	}
}
