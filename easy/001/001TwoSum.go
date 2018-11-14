package main

/**
Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
*/
func twoSum(nums []int, target int) []int {
	ret := make([]int, 0, 2)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[j] == target-nums[i] {
				ret = append(ret, i, j)
			}
		}
	}
	return ret
}

func twoSum2(nums []int, target int) []int {
	ret := make([]int, 0, 2)
	memory := make(map[int]int, 0)
	for index, num := range nums {
		res := target - num
		if _, ok := memory[res]; ok {
			ret = append(ret, memory[res], index)
			return ret
		}
		memory[num] = index

	}
	return ret
}
