package main

import "fmt"

func main() {
	fmt.Println(numUniqueEmails([]string{"test.email+alex@leetcode.com", "test.e.mail+bob.cathy@leetcode.com", "testemail+david@lee.tcode.com"}))

}

func numUniqueEmails(emails []string) int {
	res := 0
	count := make(map[string]bool)

	for _, email := range emails {
		e := ""
		after := false
		flag := false
		for _, c := range email {
			if string(c) == "+" {
				flag = true
				continue
			}

			if string(c) == "." && !after {
				continue
			}
			if string(c) == "@" {
				flag = false
				after = true
			}
			if flag {
				continue
			}
			e += string(c)
		}
		if !count[e] {
			res++
		}
		count[e] = true
	}
	for k, _ := range count {
		fmt.Println(k)
	}

	return res
}
