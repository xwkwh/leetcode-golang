package main

import "fmt"

func main() {
	fmt.Println(rob([]int{2, 3, 2}) == 3)
}

func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}

	return max(robs(nums, 1, n-1), robs(nums, 0, n-2))
}

func robs(nums []int, first, last int) int {
	pre1, pre2 := 0, 0
	for i := first; i <= last; i++ {
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
