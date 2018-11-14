package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println(reverse(123))
	fmt.Println(reverse(-123))
	fmt.Println(reverse(120))
	fmt.Println(reverse(1534236469))
	fmt.Println(reverse(-2147483648))
	fmt.Println(math.MaxInt32)
}

func reverse(x int) int {
	ret := 0
	for x != 0 {
		a := x % 10
		ret = ret*10 + a
		x = x / 10
	}
	if ret > math.MaxInt32 || ret < math.MinInt32 {
		return 0
	}
	return ret
}
