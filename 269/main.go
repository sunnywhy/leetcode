package main

import (
	"container/list"
)

// Alien Dictionary
// https://leetcode.com/problems/alien-dictionary/

func alienOrder(words []string) string {
	graph := make(map[byte]map[byte]bool)
	inDegree := make(map[byte]int)

	for _, word := range words {
		for i := 0; i < len(word); i++ {
			inDegree[word[i]] = 0
			graph[word[i]] = make(map[byte]bool)
		}
	}

	for i := 0; i < len(words)-1; i++ {
		word1, word2 := words[i], words[i+1]
		len1, len2 := len(word1), len(word2)
		if len1 > len2 && word1[:len2] == word2 {
			return ""
		}
		for j := 0; j < len1 && j < len2; j++ {
			if word1[j] == word2[j] {
				continue
			}
			if !graph[word1[j]][word2[j]] {
				graph[word1[j]][word2[j]] = true
				inDegree[word2[j]]++
			}
			break
		}
	}

	q := list.New()
	for k, v := range inDegree {
		if v == 0 {
			q.PushBack(k)
		}
	}

	var res []byte
	for q.Len() > 0 {
		e := q.Front()
		q.Remove(e)
		cur := e.Value.(byte)
		res = append(res, cur)
		for neighbor := range graph[cur] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				q.PushBack(neighbor)
			}
		}
	}
	if len(res) != len(inDegree) {
		return ""
	}
	return string(res)
}
