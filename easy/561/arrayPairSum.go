package main

import (
	"fmt"
	"sort"
)

func main() {
	A := []int {3,1,2,4}
	fmt.Println(arrayPairSum(A))

}


func arrayPairSum(nums []int) int {
	res := 0
	sort.Ints(nums)
	for i:= 0 ;i< len(nums) ;i=i+2 {
		res += nums[i]
	}
	return res
}