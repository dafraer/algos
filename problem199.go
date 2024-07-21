/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type queue struct {
    d []*TreeNode
} 

func (q *queue) Push(n *TreeNode) {
    q.d = append(q.d, n)
} 

func (q *queue) Pop() *TreeNode {
    ans := q.d[0]
    q.d = q.d[1:]
    return ans
}

func rightSideView(root *TreeNode) []int {
    var q queue
    var res []int
    q.Push(root)
    for len(q.d) != 0 {
        var rightSide *TreeNode
        qLen := len(q.d)

        for i := 0; i < qLen; i++ {
            node := q.Pop()
            if node != nil {
                rightSide = node
                q.Push(node.Left)
                q.Push(node.Right)
            }
        } 
        if rightSide != nil {
            res = append(res, rightSide.Val)
        }
    }
    return res
}
