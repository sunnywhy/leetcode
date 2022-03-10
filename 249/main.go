package main

import (
	"strings"
)

// Group Shifted Strings

func groupStrings(strings []string) [][]string {
	keyToStrings := make(map[string][]string)

	for _, s := range strings {
		key := generateKey(s)
		keyToStrings[key] = append(keyToStrings[key], s)
	}

	var res [][]string
	for _, v := range keyToStrings {
		res = append(res, v)
	}
	return res
}

func generateKey(s string) string {
	if len(s) == 1 {
		return "-"
	}
	var sb strings.Builder
	for i := 1; i < len(s); i++ {
		diff := (s[i] - s[i-1] + 26) % 26
		sb.WriteByte(diff)
		sb.WriteRune('-')
	}
	return sb.String()
}
