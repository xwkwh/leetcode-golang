package main

import "fmt"

func main() {

	fmt.Println(getRow(4))
	fmt.Println(getRow2(4))

	a := make([]int, 3)
	fmt.Println(a[2])
	fmt.Println(a[1])

	// fmt.Println(a[3])

}

func getRow(rowIndex int) []int {
	res := make([]int, 1, rowIndex+1)
	res[0] = 1
	if rowIndex == 0 {
		return res
	}
	res = append(res, 1)
	if rowIndex == 1 {
		return res
	}

	for i := 1; i < rowIndex; i++ {
		tmp := res
		res = []int{1}
		for j := 1; j < i+1; j++ {
			res = append(res, tmp[j]+tmp[j-1])
		}
		res = append(res, 1)

	}
	return res
}
func getRow2(rowIndex int) []int {
	res := make([]int, rowIndex+1)
	res[0] = 1
	for i := 1; i < len(res); i++ {
		for j := i; j >= 1; j-- {
			res[j] += res[j-1]
		}
	}
	return res
}
