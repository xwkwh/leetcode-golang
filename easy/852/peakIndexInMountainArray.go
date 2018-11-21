package main

import "fmt"

func main() {
	A := []int {0,1,2,1,0}
	fmt.Println(peakIndexInMountainArray(A))

}
func peakIndexInMountainArray(A []int) int {
	for k,a := range A {
		if k == 0 || k == len(A) -1 {
			continue
		}
		if a > A[k-1] && a > A[k+1] {
			return  k
		}
	}
	return  -1
}

func peakIndexInMountainArray2(A []int) int {
	l := 0
	r := len(A)-1
	for l < r {
		mid := (l+r) /2
		switch {
		case A[mid] < A[mid+1]:
			l = mid
		case A[mid-1] > A[mid]:
			r = mid
		default:
			return mid
		}
	}
	return  -1
}