/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Right != nil && root.Left != nil {
		return min(minDepth(root.Right), minDepth(root.Left)) + 1
	}
	if root.Right != nil {
		return minDepth(root.Right) + 1
	}
	if root.Left != nil {
		return minDepth(root.Left) + 1
	}
	return 1
}
