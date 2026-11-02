/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findMode(root *TreeNode) []int {
	mx := 0
	m := make(map[int]int)
	var f func(root *TreeNode)
	f = func(root *TreeNode) {
		if root == nil {
			return
		}
		f(root.Left)
		f(root.Right)
		m[root.Val]++
		if m[root.Val] > mx {
			mx = m[root.Val]
		}
	}
	f(root)
	res := make([]int, 0, len(m))
	for k, v := range m {
		if v == mx {
			res = append(res, k)
		}
	}
	return res
}