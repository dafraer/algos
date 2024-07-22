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

func deleteNode(root *TreeNode, key int) *TreeNode {
    if root == nil {
        return nil
    }
    if key > root.Val {
        root.Right = deleteNode(root.Right, key)
    } else if key < root.Val {
        root.Left = deleteNode(root.Left, key)
    } else {
        if root.Left == nil {
            return root.Right
        }           
        if root.Right == nil {
            return root.Left
        }
        //Find minimum form tight subtree
        cur := root.Right
        for cur.Left != nil {
            cur = cur.Left
        }
        root.Val = cur.Val
        root.Right = deleteNode(root.Right, cur.Val)
    }
    return root
}


