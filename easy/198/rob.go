package main

import "fmt"

func main() {
	fmt.Println(rob([]int{1, 2, 3, 1}) == 4)
}

func rob(nums []int) int {
	pre1, pre2 := 0, 0
	for i := 0; i < len(nums); i++ {
		cur := max(pre1, pre2+nums[i])
		pre2 = pre1
		pre1 = cur
	}
	return pre1
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
