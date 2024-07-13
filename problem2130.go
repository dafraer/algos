/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func pairSum(head *ListNode) int {
    temp := head
    //find length
    n := 0
    for temp != nil {
        n++
        temp = temp.Next
    }
    
    mx := 0
    befHead := head
    i := 0 
    for i < n/2 {
        head = head.Next
        i++
    }
    var ans *ListNode
	for head != nil {
		ans = &ListNode{head.Val, ans}
		head = head.Next
	}

    for ans != nil {
        if ans.Val+befHead.Val > mx {
            mx = ans.Val+befHead.Val
        }
        ans = ans.Next
        befHead = befHead.Next
    }
    return mx
}
