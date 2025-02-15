/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if getDepthOfUpperNode(root, p, 0) < getDepthOfUpperNode(root, q, 0) {
		for check(p, q) == false {
			p = getPrevVal(p, root, &TreeNode{})
		}
		return p
	} else {
		for check(q, p) == false {
			q = getPrevVal(q, root, &TreeNode{})
		}
		return q
	}
}

func getDepthOfUpperNode(root, t *TreeNode, ind int) int {
	if root == nil {
		return 99999999999
	}
	if root == t {
		return ind - 1
	}
	ind++
	return min(getDepthOfUpperNode(root.Left, t, ind), getDepthOfUpperNode(root.Right, t, ind))
}

func getPrevVal(t, root, prev *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root == t {
		return prev
	}
	if getPrevVal(t, root.Left, root) != nil {
		return getPrevVal(t, root.Left, root)
	} else {
		return getPrevVal(t, root.Right, root)
	}
}

// returns true if q is a subtree of p
func check(p, q *TreeNode) bool {
	if p == nil {
		return false
	}
	if p == q {
		return true
	}
	return check(p.Left, q) || check(p.Right, q)
}
