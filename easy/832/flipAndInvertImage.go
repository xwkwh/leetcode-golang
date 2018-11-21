package main

import "fmt"

func main() {
	A := [][]int{[]int{1, 1, 0}, []int{1, 0, 1}, []int{0, 0, 0}}
	fmt.Println(flipAndInvertImage(A))
}
func flipAndInvertImage(A [][]int) [][]int {
	for _, a := range A {
		for i := 0; i < len(a)/2; i ++ {
			a[i] = change(a[i])
			a[len(a)-1-i] = change(a[len(a)-1-i])
			tmp := a[i]
			a[i] = a[len(a)-1-i]
			a[len(a)-1-i] = tmp
		}
		if len(a) %2 != 0 {
			a[len(a)/2] = change(a[len(a)/2])
		}
	}
	return A
}

func change(a int) int {
	if a > 0 {
		return 0
	}
	return 1
}
