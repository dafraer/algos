/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
	flipped := flip(root)
	return areEqual(root, flipped)
}

func flip(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	newNode := &TreeNode{}
	newNode.Val = root.Val
	newNode.Right = flip(root.Left)
	newNode.Left = flip(root.Right)
	return newNode
}

func areEqual(root1, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}
	if root1 == nil || root2 == nil {
		return false
	}
	if root1.Val == root2.Val {
		return areEqual(root1.Left, root2.Left) && areEqual(root1.Right, root2.Right)
	}
	return false
}
