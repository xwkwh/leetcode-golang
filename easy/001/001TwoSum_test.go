package main

import "testing"

/**
Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
*/
func TestTwoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 13
	ret := twoSum(nums, target)
	t.Log(ret)
	ret = twoSum2(nums, target)
	t.Log(ret)

}
