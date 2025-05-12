/**␍
 * Definition for singly-linked list.␍
 * type ListNode struct {␍
 *     Val int␍
 *     Next *ListNode␍
 * }␍
 */␍
func getIntersectionNode(headA, headB *ListNode) *ListNode {␍
    m := make(map[*ListNode]struct{})␍
    hb := headB␍
    ha := headA␍
    for ha != nil {␍
        m[ha] = struct{}{}␍
        ha = ha.Next␍
    }␍
    for hb != nil {␍
        if _, ok := m[hb]; ok {␍
            return hb␍
        }␍
        hb = hb.Next␍
    }␍
    return nil␍
}
