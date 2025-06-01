	/**
	 * Definition for a binary tree node.
	 * type TreeNode struct {
	 *     Val int
	 *     Left *TreeNode
	 *     Right *TreeNode
	 * }
	 */
	func binaryTreePaths(root *TreeNode) []string {
		ans := make([]string, 0)
		var dfs func(b []byte, root *TreeNode)
		dfs = func(b []byte, root *TreeNode) {
			if root == nil {
				ans = append(ans, string(b[:len(b)-2]))
				return
			}
			b = append(b, []byte(strconv.Itoa(root.Val))...)
			b = append(b, '-')
			b = append(b, '>')
			if root.Right != nil {
				x := make([]byte, len(b))
				copy(x, b)
				dfs(x, root.Right)
			}
			if root.Left != nil {
				x := make([]byte, len(b))
				copy(x, b)
				dfs(x, root.Left)
			}
			if root.Right == nil && root.Left == nil {
				ans = append(ans, string(b[:len(b)-2]))
				return
			}
		}
		b := make([]byte, 0, 100)
		dfs(b, root)
		return ans
	}
