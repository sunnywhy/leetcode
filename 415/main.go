package main

func addStrings(num1 string, num2 string) string {
	if len(num1) < len(num2) {
		return addStrings(num2, num1)
	}
	var carry int
	k := len(num1)
	res := make([]byte, k+1)
	i, j := len(num1)-1, len(num2)-1

	for i >= 0 || carry > 0 {
		sum := carry
		if i >= 0 {
			sum += int(num1[i] - '0')
			i--
		}
		if j >= 0 {
			sum += int(num2[j] - '0')
			j--
		}
		res[k] = byte(sum%10 + '0')
		k--
		carry = sum / 10
	}
	return string(res[k+1:])
}
