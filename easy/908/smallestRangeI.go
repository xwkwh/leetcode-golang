// Package main
package main

import "fmt"

func main() {
	A := []int{1, 3, 6}
	K := 3
	fmt.Println(smallestRangeI(A, K))
}

func smallestRangeI(A []int, K int) int {
	i := A[0]
	j := A[0]
	for _, a := range A {
		i = min(i, a)
		j = max(j, a)
	}
	return max(0, j-K-i-K)
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
