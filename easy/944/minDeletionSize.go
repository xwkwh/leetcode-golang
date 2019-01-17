package main

import (
	"fmt"
	"reflect"
)

func main() {
	A := []string {"cba","daf","ghi"}
	fmt.Println(minDeletionSize(A))
	A = []string {"a","b"}
	fmt.Println(minDeletionSize(A))
	A = []string {"rrjk","furt","guzm"}
	fmt.Println(minDeletionSize(A))

	a := "asc"
	fmt.Println(a[0],"-",a[2])
	fmt.Println(reflect.TypeOf(a[0]))
}
func minDeletionSize(A []string) int {
	res := 0
	if len(A) == 0 {return 0}
	l := len(A)
	for j:=0 ;j< len(A[0]);j++ {
		for i:= 1 ; i< l ;i ++ {
			if A[i][j] < A[i-1][j] {
				res ++
				break
			}
		}
	}
	return res
}
