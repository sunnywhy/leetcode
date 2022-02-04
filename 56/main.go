package main

import "sort"
// Merge Intervals
func merge(intervals [][]int) [][]int {
    if len(intervals) <= 1 {
        return intervals
    }
    // sort
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })
    prevEnd := intervals[0][1]
    res := [][]int{}
    res = append(res, intervals[0])
    
    for i := 1; i < len(intervals); i++ {
        if intervals[i][0] > prevEnd {
            res[len(res)-1][1] = prevEnd
            res = append(res, intervals[i])
            prevEnd = intervals[i][1]
        }else {
            if intervals[i][1] > prevEnd {
                prevEnd = intervals[i][1]
            }
        }
    }
    res[len(res)-1][1] = prevEnd
    return res
}