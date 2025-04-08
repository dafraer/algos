/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	var m [][]int
	bfs(root, &m, 0)
	return m
}

func bfs(root *TreeNode, m *[][]int, level int) {
	if root == nil {
		return
	}
	for len(*m) < level+1 {
		*m = append(*m, []int{})
	}
	(*m)[level] = append((*m)[level], root.Val)
	bfs(root.Left, m, level+1)
	bfs(root.Right, m, level+1)
}
