package main

import (
	"strconv"
	"strings"
)

// Exclusive Time of Functions

func exclusiveTime(n int, logs []string) []int {
	res := make([]int, n)
	var prevTime int
	var stack []int
	for _, log := range logs {
		data := strings.Split(log, ":")
		curId, _ := strconv.Atoi(data[0])
		curTime, _ := strconv.Atoi(data[2])

		if data[1] == "start" {
			if len(stack) > 0 {
				prevId := stack[len(stack)-1]
				res[prevId] += curTime - prevTime
			}
			prevTime = curTime
			stack = append(stack, curId)
		} else {
			curTime = curTime + 1
			res[curId] += curTime - prevTime
			stack = stack[:len(stack)-1]
			prevTime = curTime
		}
	}
	return res
}
