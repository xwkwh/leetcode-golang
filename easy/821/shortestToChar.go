package main

import "fmt"

func main() {
	S := "loveleetcode"
	C := byte('e')
	fmt.Println(shortestToChar(S, C))
}

func shortestToChar(S string, C byte) []int {

	l := len(S)
	res := make([]int, l)
	for i := 0; i < l; i++ {
		left, right := 0, 0
		for j := i; j >= 0; j-- {
			if S[j] == C {
				break
			}
			left++
			if j == 0 {
				left = l
			}
		}
		for k := i; k < l; k++ {
			if S[k] == C {
				break
			}
			right++
			if k == l-1 {
				right = l
			}
		}
		if left > right {
			res[i] = right
		} else {
			res[i] = left
		}
	}
	return res
}
