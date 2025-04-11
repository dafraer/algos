/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	return dfs(root, math.MinInt, math.MaxInt)
}

func dfs(root *TreeNode, mn, mx int) bool {
	if root == nil {
		return true
	}
	if root.Val <= mn || root.Val >= mx {
		return false
	}
	return dfs(root.Left, mn, root.Val) && dfs(root.Right, root.Val, mx)
}   

