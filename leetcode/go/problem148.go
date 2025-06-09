	/**
	 * Definition for singly-linked list.
	 * type ListNode struct {
	 *     Val int
	 *     Next *ListNode
	 * }
	 */
	func sortList(head *ListNode) *ListNode {
		//determine size
		h := head
		size := 0
		for h != nil {
			size++
			h = h.Next
		}

		var mergeSort func(head *ListNode, size int) *ListNode
		mergeSort = func(head *ListNode, size int) *ListNode {
			if head == nil {
				return nil
			}
			if size == 1 {
				return &ListNode{Val: head.Val, Next: nil}
			}
			m := size / 2
			h := head
			cnt := 0
			for cnt < m && h != nil {
				h = h.Next
				cnt++
			}
			srt1 := mergeSort(head, m)
			srt2 := mergeSort(h, size-m)
			if srt1.Val > srt2.Val {
				srt1, srt2 = srt2, srt1
			}
			ans := &ListNode{}
			ansIt := ans
			for srt1 != nil || srt2 != nil {
				if srt1 == nil {
					ansIt.Val = srt2.Val
					srt2 = srt2.Next
				} else if srt2 == nil {
					ansIt.Val = srt1.Val
					srt1 = srt1.Next
				} else if srt1.Val > srt2.Val {
					ansIt.Val = srt2.Val
					srt2 = srt2.Next
				} else {
					ansIt.Val = srt1.Val
					srt1 = srt1.Next
				}
				if srt1 != nil || srt2 != nil {
					ansIt.Next = &ListNode{}
					ansIt = ansIt.Next
				}
			}
			return ans
		}
		return mergeSort(head, size)
	}
