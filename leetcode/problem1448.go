/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func goodNodes(root *TreeNode) int {
	return count(root, root.Val) + 1
}

func count(root *TreeNode, mx int) int {
	mx = max(mx, root.Val)
	switch {
	case root.Right != nil && root.Left != nil:
		if mx <= root.Left.Val && mx <= root.Right.Val {
			return count(root.Left, mx) + count(root.Right, mx) + 2
		}
		if mx <= root.Right.Val {
			return count(root.Left, mx) + count(root.Right, mx) + 1
		}
		if mx <= root.Left.Val {
			return count(root.Left, mx) + count(root.Right, mx) + 1
		}
		return count(root.Left, mx) + count(root.Right, mx)
	case root.Left != nil:
		if mx <= root.Left.Val {
			return count(root.Left, mx) + 1
		}
		return count(root.Left, mx)
	case root.Right != nil:
		if mx <= root.Right.Val {
			return count(root.Right, mx) + 1
		}
		return count(root.Right, mx)
	default:
		return 0
	}
}
