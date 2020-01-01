package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(largestPerimeter([]int{3,2,3,4}))
}

func largestPerimeter(A []int) int {
	sort.Ints(A)

	for i := len(A) -1 ; i > 1 ; i -- {
		if A[i ] < A[i-1] + A[i-2] {
			return  A[i] + A[i-1] + A[i-2]
		}
	}
	return  0
}