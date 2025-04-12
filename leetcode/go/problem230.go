/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
	s := make([]int, 0)
	add(root, &s)
	fmt.Println(s)
	for i := 0; i < len(s); i++ {
		if i == k-1 {
			return s[i]
		}
	}
	return s[0]
}

func add(root *TreeNode, s *[]int) {
	if root.Left != nil {
		add(root.Left, s)
	}
	*s = append(*s, root.Val)
	if root.Right != nil {
		add(root.Right, s)
	}
}
