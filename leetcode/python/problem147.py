# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def insertionSortList(self, head: Optional[ListNode]) -> Optional[ListNode]:
        a = []
        while head:
            a.append(head.val)
            head = head.next
        a.sort()
        ans = ListNode()
        h = ans
        for i in range(len(a)):
            h.val = a[i]
            if i < len(a) - 1:
                h.next = ListNode()
                h = h.next
        return ans
