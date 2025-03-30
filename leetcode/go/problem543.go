/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var res int

func diameterOfBinaryTree(root *TreeNode) int {
	res = 0
	_ = dp(root)
	return res
}

func dp(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := dp(root.Left)
	r := dp(root.Right)
	res = max(res, l+r)
	return max(l, r) + 1
}
