package main

import (
	"fmt"
)

func main() {
	A := "IDID"
	fmt.Println(diStringMatch(A))
	A = "III"
	fmt.Println(diStringMatch(A))
	A = "DDI"
	fmt.Println(diStringMatch(A))

	fmt.Println()
}

func diStringMatch(S string) []int {

	start := 0
	end := len(S)
	res := []int{}
	for _, c := range []rune(S) {
		if rune('D') == c {
			res = append(res, end)
			end--
		} else {
			res = append(res, start)
			start++
		}
	}
	res = append(res, start)
	return res
}
