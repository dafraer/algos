/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumOfLeftLeaves(root *TreeNode) int {
	return sum(root, false)
}

func sum(root *TreeNode, left bool) int {
	if root == nil {
		return 0
	}
	if left && root.Left == nil && root.Right == nil {
		return sum(root.Left, true) + sum(root.Right, false) + root.Val
	}
	return sum(root.Left, true) + sum(root.Right, false)
}