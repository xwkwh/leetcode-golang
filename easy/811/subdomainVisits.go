package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	in := []string{"9001 discuss.leetcode.com"}

	fmt.Println(subdomainVisits(in))

	in = []string{"900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"}

	fmt.Println(subdomainVisits(in))
}

func subdomainVisits(cpdomains []string) []string {
	mem := make(map[string]int, 0)

	for _, cpdomain := range cpdomains {
		ss := strings.Split(cpdomain, " ")
		// doain := ss[1]
		domains := strings.Split(ss[1], ".")
		count, _ := strconv.Atoi(ss[0])
		l := len(domains)
		for i := 0; i < l; i++ {
			mem[strings.Join(domains[i:], ".")] += count
		}
	}
	res := make([]string, 0)
	for k, v := range mem {
		str := fmt.Sprintf("%d %s", v, k)
		res = append(res, str)
	}
	return res
}

func subdomainVisits(cpdomains []string) []string {
	times := make(map[string]int, 0)
	var out []string
	for _, pair := range cpdomains {
		splited := strings.Split(pair, " ")
		count, _ := strconv.Atoi(splited[0])
		domain := splited[1]
		times[domain] += count
		for index := strings.Index(domain, "."); index != -1; index = strings.Index(domain, ".") {
			domain = domain[index+1:]
			times[domain] += count
		}
	}
	for domain, count := range times {
		out = append(out, fmt.Sprintf("%d %s", count, domain))
	}
	return out
}
