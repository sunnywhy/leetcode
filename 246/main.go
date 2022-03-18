package main

// Strobogrammatic Number

func isStrobogrammatic(num string) bool {
	mapping := map[byte]byte{'0': '0', '1': '1', '6': '9', '8': '8', '9': '6'}

	for i, j := 0, len(num)-1; i <= j; i, j = i+1, j-1 {
		if mapping[num[i]] != num[j] {
			return false
		}
	}
	return true
}
