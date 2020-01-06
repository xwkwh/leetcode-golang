package main

func numberOfArithmeticSlices(A []int) int {
	if len(A) <= 2 {
		return 0
	}
	dp := make([]int, len(A))
	for i := 2; i < len(A); i++ {
		if A[i]-A[i-1] == A[i-1]-A[i-2] {
			dp[i] = dp[i-1] + 1
		}
	}

	total := 0
	for _, v := range dp {
		total += v
	}
	return total

}
