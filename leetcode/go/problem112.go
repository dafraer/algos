/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	return dfs(root, targetSum, 0)
}

func dfs(root *TreeNode, targetSum int, cur int) bool {
	if root.Right == nil && root.Left == nil {
		return targetSum == cur+root.Val
	}
	if root.Right == nil {
		return dfs(root.Left, targetSum, cur+root.Val)
	}
	if root.Left == nil {
		return dfs(root.Right, targetSum, cur+root.Val)
	}
	return dfs(root.Left, targetSum, cur+root.Val) || dfs(root.Right, targetSum, cur+root.Val)

}
