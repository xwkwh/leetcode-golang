package main

import (
	"fmt"
	"sort"
)

func main() {
	input := []int{-4, -1, 0, 3, 10}
	fmt.Println(sortedSquares(input))
}

func sortedSquares(A []int) []int {
	for i, a := range A {
		A[i] = a * a
	}
	sort.Ints(A)
	return A
}
