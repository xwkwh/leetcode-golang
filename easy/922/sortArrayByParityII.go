// Package main provides ...
package main

import "fmt"

func main() {
	A := []int{3, 1, 4, 2}
	res := sortArrayByParityII(A)
	fmt.Println(res)
}

func sortArrayByParityII(A []int) []int {
	// res := [,{]int{}
	odd := 1
	for i := 0; i < len(A)-1; i += 2 {
		if A[i]%2 != 0 {
			for j := odd; j < len(A); j += 2 {
				if A[j]%2 == 0 {
					odd = j + 2

					tmp := A[j]
					A[j] = A[i]
					A[i] = tmp

					break
				}
			}
		}
	}
	return A
}
