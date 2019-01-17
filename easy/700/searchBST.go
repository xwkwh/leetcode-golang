package main

func main() {

}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	// if root.val == val {
	// return root
	// }

	switch {
	case root.Val == val:
		return root
	case root.Val < val:
		return searchBST(root.Right, val)
	case root.Val > val:
		return searchBST(root.Left, val)
	}
	return nil
}
