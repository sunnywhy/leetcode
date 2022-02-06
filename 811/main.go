package main

import (
	"strconv"
	"strings"
)

// Subdomain Visit Count
func subdomainVisits(cpdomains []string) []string {
	data := make(map[string]int)
	for _, line := range cpdomains {
		lineArray := strings.Split(line, " ")
		count, _ := strconv.Atoi(lineArray[0])
		domain := lineArray[1]

		data[domain] += count

		for strings.Index(domain, ".") != -1 {
			domain = string([]rune(domain)[strings.Index(domain, ".")+1:])
			data[domain] += count
		}
	}

	res := []string{}
	for k, v := range data {
		res = append(res, strconv.Itoa(v)+" "+k)
	}
	return res
}

func main() {
	cpdomains := []string{"900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"}
	for _, v := range subdomainVisits(cpdomains) {
		println(v)
	}
}
