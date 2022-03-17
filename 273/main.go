package main

import "strings"

//  Integer to English Words

var LESSTHANTWENTIES = []string{"", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten",
	"Eleven", "Twelve", "Thirteen", "Fourteen", "Fifteen", "Sixteen", "Seventeen", "Eighteen", "Nineteen"}
var TENS = []string{"", "Ten", "Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy", "Eighty", "Ninety"}
var THOUSANDS = []string{"", "Thousand", "Million", "Billion"}

func numberToWords(num int) string {
	if num == 0 {
		return "Zero"
	}
	var res string
	var count int
	for num > 0 {
		if num%1000 != 0 {
			res = lessThanThousand(num%1000) + THOUSANDS[count] + " " + res
		}
		count++
		num /= 1000
	}
	return strings.TrimSpace(res)
}

func lessThanThousand(num int) string {
	if num == 0 {
		return ""
	}
	if num < 20 {
		return LESSTHANTWENTIES[num] + " "
	}

	if num < 100 {
		return TENS[num/10] + " " + lessThanThousand(num%10)
	}

	return LESSTHANTWENTIES[num/100] + " Hundred " + lessThanThousand(num%100)
}
