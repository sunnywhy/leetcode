package main

import (
	"math"
	"strings"
)

func addBoldTag(s string, words []string) string {
	var intervals [][]int
	chars := []byte(s)

	for i := 0; i < len(chars); i++ {
		for _, w := range words {
			if strings.HasPrefix(s[i:], w) {
				intervals = append(intervals, []int{i, i + len(w) - 1})
			}
		}
	}

	var merged [][]int
	for _, v := range intervals {
		if len(merged) == 0 || merged[len(merged)-1][1] < v[0]-1 {
			merged = append(merged, []int{v[0], v[1]})
		} else {
			merged[len(merged)-1][1] = int(math.Max(float64(v[1]), float64(merged[len(merged)-1][1])))
		}
	}

	var prev int
	var sb strings.Builder
	for _, v := range merged {
		sb.Write(chars[prev:v[0]])
		sb.WriteString("<b>")
		sb.Write(chars[v[0] : v[1]+1])
		sb.WriteString("</b>")
		prev = v[1] + 1
	}
	sb.Write(chars[prev:])
	return sb.String()
}
