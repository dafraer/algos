/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	h := head
	n := &Node{}
	ans := n
	for h != nil {
		n.Val = h.Val
		n.Next = &Node{}
		n.Random = h.Random
		h = h.Next
		if h != nil {
			n = n.Next
		}
	}
	n.Next = nil
	n = ans
	for n != nil {
		h = head
		newN := ans
		for h != nil {
			if h == n.Random {
				n.Random = newN
			}
			newN = newN.Next
			h = h.Next
		}
		n = n.Next
	}
	return ans
}
