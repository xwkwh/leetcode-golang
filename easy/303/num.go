package main

type NumArray struct {
	Num []int
}

func Constructor(nums []int) NumArray {
	sum := make([]int, len(nums)+1)
	for i := 1; i < len(sum); i++ {
		sum[i] = sum[i-1] + nums[i-1]
	}
	return NumArray{sum}
}

func (this *NumArray) SumRange(i int, j int) int {
	return this.Num[j+1] - this.Num[i]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */
