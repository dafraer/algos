/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	return reflect.DeepEqual(getLVS(root1), getLVS(root2))
}

func getLVS(root *TreeNode) []int {
	s := make([]int, 0)
	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}
	if root.Left != nil {
		s = append(s, getLVS(root.Left)...)
	}
	if root.Right != nil {
		s = append(s, getLVS(root.Right)...)
	}
	return s
}
