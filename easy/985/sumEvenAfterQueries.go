package main

import "fmt"

func main() {
	A := []int{1, 2, 3, 4}
	queries := [][]int{
		[]int{1, 0}, []int{-3, 1}, []int{-4, 0}, []int{2, 3},
	}
	fmt.Println(sumEvenAfterQueries(A, queries))
}

func sumEvenAfterQueries(A []int, queries [][]int) []int {
	res := make([]int, len(A))
	for k, v := range queries {
		A[v[1]] += v[0]
		res[k] = getEvenSum(A)
	}
	return res
}

func getEvenSum(A []int) int {
	res := 0
	for _, v := range A {
		if v%2 == 0 {
			res += v
		}
	}
	return res
}
