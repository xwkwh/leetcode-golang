package main

import (
	"fmt"
	"math"
)

func judgeSquareSum(c int) bool {
	if c < 0 {
		return false
	}

	i, j := 0, int(math.Sqrt(float64(c)))
	for i < j {
		sum := i*i + j*j
		if sum < c {
			i++
		} else if sum > c {
			j--
		} else {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(judgeSquareSum(5))
	fmt.Println(judgeSquareSum(10))
	fmt.Println(judgeSquareSum(9))
	fmt.Println(judgeSquareSum(3))
	fmt.Println(math.Sqrt(5))
	fmt.Println(math.Sqrt(10))
}
