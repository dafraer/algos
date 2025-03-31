/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if math.Abs(float64(depth(root.Right)-depth(root.Left))) > 1 {
		return false
	}
	return isBalanced(root.Right) && isBalanced(root.Left)
}

func depth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(depth(root.Right), depth(root.Left)) + 1
}
