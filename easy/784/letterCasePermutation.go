package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println(letterCasePermutation("a1b2"))
}

func letterCasePermutation(S string) []string {

	l := len(S)
	var res []string
	for i := 0; i < l; i++ {
		part := S[:i]
		if strings.ToLower(part) != part {
			res = append(res, strings.ToLower(part)+S[i:])
			continue
		}

		if strings.ToUpper(part) != part {
			res = append(res, strings.ToUpper(part)+S[i:])
			continue
		}
	}
	res = append(res, S)
	return res
}
