package main

// Valid Number

func isNumber(s string) bool {
	var numberSeen, dotSeen, eSeen bool

	for i, c := range s {
		if c >= '0' && c <= '9' {
			numberSeen = true
		} else if c == '.' {
			if dotSeen || eSeen {
				return false
			}
			dotSeen = true
		} else if c == 'e' || c == 'E' {
			if eSeen || !numberSeen {
				return false
			}
			eSeen = true
			numberSeen = false
		} else if c == '+' || c == '-' {
			if numberSeen {
				return false
			}
			if i > 0 && s[i-1] == '.' {
				return false
			}
		} else {
			return false
		}
	}
	return numberSeen
}
