	/**
	 * Definition for singly-linked list.
	 * type ListNode struct {
	 *     Val int
	 *     Next *ListNode
	 * }
	 */
	func getDecimalValue(head *ListNode) int {
		b := strings.Builder{}
		for head != nil {
			b.WriteString(strconv.Itoa(head.Val))
			head = head.Next
		}
		n, _ := strconv.ParseInt(b.String(), 2, 64)
		return int(n)
	}
