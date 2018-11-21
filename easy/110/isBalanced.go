package main

import "math"

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {

	if nil == root {
		return true
	}

	left := depth(root.Left)
	right := depth(root.Right)

	return math.Abs(float64(left-right)) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)

}

func depth(root *TreeNode) int {
	if nil == root {
		return 0
	}
	return int(math.Max(float64(depth(root.Left), float64(depth(root.Right))))) + 1
}

func isBalanced2(root *TreeNode) bool {
	return dfs(root) != -1
}

func dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := dfs(root.Left)
	right := dfs(root.Right)

	if abs(left-right) > 1 || left == -1 || right == -1 {
		return -1
	}

	return max(left, right) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return 0 - a
}
