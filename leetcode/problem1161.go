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

func maxLevelSum(root *TreeNode) int {
    var q queue
    q.Push(root)
    mx := root.Val 
    mxInd := 1
    ind := 1
    for len(q.d) > 0 {
        sum := 0
        check := false
        qLen := len(q.d)

        for i := 0; i < qLen; i++ {
            curNode := q.Pop()
            if curNode != nil {
                q.Push(curNode.Left)
                q.Push(curNode.Right)
                check = true
                sum += curNode.Val
            }
        }
        if sum > mx && check {
            mx = sum
            mxInd = ind
        }
        ind++
    }
    return mxInd
}
