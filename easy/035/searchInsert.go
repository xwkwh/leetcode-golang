package main

import "fmt"

func main() {
	nums := []int{1, 3, 5, 6}
	fmt.Println(searchInsert(nums, 5))
	fmt.Println(searchInsert(nums, 2))
	fmt.Println(searchInsert(nums, 7))
	fmt.Println(searchInsert(nums, 0))
}

/**
Given a sorted array and a target value, return the index if the target is found.
If not, return the index where it would be if it were inserted in order.
*/
func searchInsert(nums []int, target int) int {
	for index, num := range nums {
		if target <= num {
			return index
		}
	}
	return len(nums)
}
