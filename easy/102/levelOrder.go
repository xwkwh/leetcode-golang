package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	ret := make([][]int, 0)

	var dfs func(*TreeNode, int)

	dfs := func(root *TreeNode, level int) {
		if root == nil {
			return
		}
		//有新的level，则先分配
		if level >= len(ret) {
			ret = append(ret, []int{})
		}
		ret[level] = append(ret[level], root.Val)
		dfs(root.Left, level+1)
		dfs(root.Right, level+1)
	}

	dfs(root, 0)

	return ret

}
func levelOrder2(root *TreeNode) [][]int {
	ret := make([][]int)

	var dfs func(*TreeNode, int)

	dfs := func(root *TreeNode, level int) {
		if root == nil {
			return
		}
		//有新的level，则先分配
		if level >= len(ret) {
			ret = append(ret, []int{})
		}
		ret[level] = append(ret[level], root.Val)
		dfs(root.Left, level+1)
		dfs(root.Right, level+1)
	}

	dfs(root, 0)

	return ret

}
