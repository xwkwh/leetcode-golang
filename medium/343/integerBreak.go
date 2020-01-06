package main

import "fmt"

func main() {
	fmt.Println(aintegerBreak(2) == 1)
	fmt.Println(aintegerBreak(10) == 36)
}
func aintegerBreak(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= i-1; j++ {
			cur := max(dp[i-j]*j, j*(i-j))
			dp[i] = max(cur, dp[i])
		}
	}
	return dp[n]
}
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
