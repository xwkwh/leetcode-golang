package main

import "fmt"

func main() {
	var arr []int = []int{1, 0, 2, 3, 0, 4, 5, 0}
	duplicateZeros(arr)

	fmt.Println(arr)
}

func duplicateZeros(arr []int) {
	var res []int
	for _, v := range arr {
		res = append(res, v)
		if v == 0 {
			res = append(res, 0)
		}
	}
	for k := range arr {
		arr[k] = res[k]
	}

	fmt.Println(arr)
}
