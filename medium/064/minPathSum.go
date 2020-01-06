package main

func minPathSum(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	m := len(grid)
	n := len(grid[0])
	dp := make([]int, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if j == 0 {
				dp[j] = dp[j] // 只能从上侧走到该位置
			} else if i == 0 {
				dp[j] = dp[j-1] // 只能从左侧走到该位置
			} else {
				dp[j] = min(dp[j-1], dp[j]) // 是从上往下走还是从左往右走
			}
			dp[j] += grid[i][j]
		}
	}

	return dp[n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minPathSum2(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j != 0 {
				grid[i][j] += grid[0][j-1]
			}
			if i != 0 && j == 0 {
				grid[i][j] += grid[i-1][0]
			}
			if i != 0 && j != 0 {
				grid[i][j] += min(grid[i][j-1], grid[i-1][j])
			}
		}
	}
	return grid[i-1][j-1]
}
