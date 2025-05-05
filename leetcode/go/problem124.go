/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum(root *TreeNode) int {
	mx, _ := mxp(root)
	return mx
}

func mxp(root *TreeNode) (mx int, mxExit int) {
	if root.Right != nil && root.Left != nil {
		mxRight, mxExitRight := mxp(root.Right)
		mxLeft, mxExitLeft := mxp(root.Left)
		mx = max(mxExitRight+mxExitLeft+root.Val, max(mxRight, mxLeft))
		mxExit = max(0, max(mxExitLeft+root.Val, mxExitRight+root.Val))
		return
	}
	if root.Right != nil {
		mxRight, mxExitRight := mxp(root.Right)
		mx = max(mxExitRight+root.Val, mxRight)
		mxExit = max(mxExitRight+root.Val, 0)
		return
	}
	if root.Left != nil {
		mxLeft, mxExitLeft := mxp(root.Left)
		mx = max(mxExitLeft+root.Val, mxLeft)
		mxExit = max(mxExitLeft+root.Val, 0)
		return
	}
	return root.Val, max(root.Val, 0)
}
