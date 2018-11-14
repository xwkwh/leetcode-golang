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
func levelOrderBottom(root *TreeNode) [][]int {
	ret := make([][]int, 0)

	var dfs func(*TreeNode, int)

	dfs := func(root *TreeNode, level int) {
		if root == nil {
			return
		}
		//有新的level，则先分配
		if level >= len(ret) {
			ret = append([][]int{[]int{}}, ret...)
		}
		ret[len(ret)-level-1] = append(ret[len(ret)-level-1], root.Val)
		dfs(root.Left, level+1)
		dfs(root.Right, level+1)
	}

	dfs(root, 0)

	return ret

}

func solution(root *TreeNode) [][]int {
	ret := make([][]int, 0)
	if nil == root {
		return ret
	}

	next := []*TreeNode{root}

	level(next)

	return ret
}

func level([list]*TreeNode) {
	if 0 == len(list) {
		return
	}
}
