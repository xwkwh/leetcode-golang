package main

var direct = [][]int{
	[]int{0, 1},
	[]int{0, -1},
	[]int{1, 0},
	[]int{-1, 0},
}
var m, n int

func maxAreaOfIsland(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	m = len(grid)
	n = len(grid[0])
	maxArea := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			maxArea = max(maxArea, dfs(grid, i, j))
		}
	}
	return maxArea
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func dfs(grid [][]int, i, j int) int {

	// 1. 异常终止条件 (越界 or 走过了)
	if i < 0 || j < 0 || i >= m || j >= n {
		return 0
	}
	if grid[i][j] == 0 {
		return 0
	}

	// 2. 没有异常 是不是得算这一步
	grid[i][j] = 0
	area := 1
	// 3. 走完这一步呢 继续走四个方向
	for _, v := range direct {
		area += dfs(grid, i+v[0], j+v[1])
	}

	return area
}
