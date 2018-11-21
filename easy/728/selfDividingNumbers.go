package main

import (
	"fmt"
)

func main() {
	fmt.Println()
	fmt.Println(selfDividingNumbers(66, 708))
}

func selfDividingNumbers(left int, right int) []int {
	res := []int{}
	for i := left; i <= right; i ++ {
		if i < 10 {
			res = append(res, i)
			continue
		}
		if i%10 == 0 {
			continue
		}
		is := true
		for j := i; j > 0; j = j / 10 {
			if j%10 == 0 {
				is = false
				continue
			}
			k := j % 10
			if i%k != 0 {
				is = false
				continue
			}
		}
		if is {
			res = append(res, i)
		}
	}
	return res
}
