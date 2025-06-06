	/**
	 * Definition for a binary tree node.
	 * type TreeNode struct {
	 *     Val int
	 *     Left *TreeNode
	 *     Right *TreeNode
	 * }
	 */
	type zigZagNode struct {
		node  *TreeNode
		level int
	}

	type Queue []*zigZagNode

	func newQueue() *Queue {
		s := Queue(make([]*zigZagNode, 0))
		return &s
	}

	func (q *Queue) Push(node *zigZagNode) {
		*q = append(*q, node)
	}

	func (q *Queue) Pop() *zigZagNode {
		if len([]*zigZagNode(*q)) == 0 {
			return nil
		}
		ans := (*q)[0]
		(*q)[0] = nil
		*q = (*q)[1:]
		return ans
	}

	func zigzagLevelOrder(root *TreeNode) [][]int {
		if root == nil {
			return nil
		}
		q := newQueue()
		q.Push(&zigZagNode{root, 0})
		ans := make([][]int, 1)
		for n := q.Pop(); n != nil; n = q.Pop() {
			if n.node == nil {
				continue
			}
			q.Push(&zigZagNode{n.node.Right, n.level + 1})
			q.Push(&zigZagNode{n.node.Left, n.level + 1})
			if n.level >= len(ans) {
				ans = append(ans, []int{n.node.Val})
			} else {
				ans[n.level] = append(ans[n.level], n.node.Val)
			}
		}
		for i := 0; i < len(ans); i++ {
			if i%2 == 0 {
				slices.Reverse(ans[i])
			}
		}
		return ans
	}
