package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(bitwiseComplement(5) == 2)
	fmt.Println(bitwiseComplement(7) == 0)
	fmt.Println(bitwiseComplement(10) == 5)
	fmt.Println(bitwiseComplement(0) == 1)
}

func bitwiseComplement(N int) int {
	if N == 0 {
		return 1
	}
	i := 0
	sum := 0
	for N != 0 {

		bit := N & 1
		if bit == 0 {
			bit = 1
		} else {
			bit = 0
		}
		sum += int(math.Pow(2, float64(i))) * bit
		i++
		N = N >> 1

	}
	return sum
}
