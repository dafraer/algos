/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func pushBack(n *ListNode, val int) *ListNode {
	if n == nil {
		return &ListNode{val, nil}
	}
	n.Next = &ListNode{val, nil}
	return n.Next
}

func makeRotate(val int) func() int {
	i := -1
	return func() int {
		if i >= val-1 {
			i = 0
		} else {
			i++
		}
		return i
	}
}

func mergeKLists(lists []*ListNode) *ListNode {
	ans := &ListNode{}
	ansFr := ans
	rotate := makeRotate(len(lists))
	i := rotate()
	cntNil := 0
	m := make(map[int]struct{})
	mn := math.MaxInt
	mnInd := 0
	for cntNil < len(lists) {
		if lists[i] == nil {
			if _, ok := m[i]; !ok {
				m[i] = struct{}{}
				cntNil++
			}
		} else {
			if lists[i].Val < mn {
				mn = lists[i].Val
				mnInd = i
			}
		}
		i = rotate()
		if i == 0 {
			if mn != math.MaxInt {
				ans = pushBack(ans, mn)
				lists[mnInd] = lists[mnInd].Next
			}
			mn = math.MaxInt
		}
	}
	return ansFr.Next
}
