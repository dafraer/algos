/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	s := make([]int, 0, 100)
	bfs(root, &s)
	return s
}

func bfs(root *TreeNode, s *[]int) {
	if root == nil {
		return
	}
	bfs(root.Left, s)
	*s = append(*s, root.Val)
	bfs(root.Right, s)
}
