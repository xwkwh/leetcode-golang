package main

import "fmt"

func main() {
	A := []int {3,1,2,4}
	fmt.Println(sortArrayByParity(A))

}
func sortArrayByParity(A []int) []int {
	even := 0
	for i := 0 ; i < len(A) ; i++ {
		if A[i] %2 == 0 {
			tmp := A[even]
			A[even] = A[i]
			A[i] = tmp
			even ++
		}
	}
	return A
}
