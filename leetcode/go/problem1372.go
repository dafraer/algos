/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func longestZigZag(root *TreeNode) int {

    return max(count(root.Left, -1, 0, 0), count(root.Right, 1, 0, 0))
}

func count(root *TreeNode, dir int, cur int, mx int) int {
    if cur > mx {
        mx = cur
    }
    if root == nil {
        return mx
    }
    if dir == -1 {
        return max(count(root.Left, -1, 0, mx), count(root.Right, 1, cur+1, mx))
    } else {
        return max(count(root.Left, -1, cur+1, mx), count(root.Right, 1, 0, mx))
    }
}

