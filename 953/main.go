package main

// Verifying an Alien Dictionary

func isAlienSorted(words []string, order string) bool {
	var orderDict [26]int
	for i, c := range order {
		orderDict[int(c-'a')] = i
	}

	for i := 0; i < len(words)-1; i++ {
		if !valid(words[i], words[i+1], orderDict) {
			return false
		}
	}
	return true
}

func valid(w1, w2 string, orderDict [26]int) bool {
	m, n := len(w1), len(w2)
	var i int
	for i < m && i < n {
		if w1[i] == w2[i] {
			i++
		} else {
			return orderDict[int(w1[i]-'a')] < orderDict[int(w2[i]-'a')]
		}
	}
	return i == m
}
