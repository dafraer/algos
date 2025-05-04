/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) < 1 {
		return nil
	}
	var root = &TreeNode{}
	m := (len(nums) - 1) / 2
	root.Val = nums[m]
	root.Left = sortedArrayToBST(nums[:m])
	root.Right = sortedArrayToBST(nums[m+1:])
	return root
}
