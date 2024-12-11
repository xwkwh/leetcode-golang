package main

import "fmt"

/**
Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
*/

func main() {
	fmt.Println(twoSum2022([]int{2, 7, 11, 15}, 9))

}

// 没注意题目呀, 只出现一次但是不一定有序
func twoSum2022(nums []int, target int) []int {
	left, right := 0, len(nums)-1
	res := make([]int, 2)
	for left < right {
		sum := nums[left] + nums[right]
		if sum == target {
			res[0], res[1] = left, right
			break
		}
		if sum < target {
			left++
		}
		if sum > target {
			right--
		}
	}
	return res
}

// 从左向右双重循环遍历
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

// 利用只出现一次的特征使用map存储element对应的index，单次遍历即可解决
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

// 错误的解法
func twoSum3(nums []int, targer int) []int {
	ret := make([]int, 0, 2)
	i, j := 0, len(nums)-1
	for i <= j {
		sum := nums[i] + nums[j]
		if sum < targer {
			i++
		} else if sum > targer {
			j--
		} else {
			ret = append(ret, i, j)
			return ret
		}

	}
	return ret
}
