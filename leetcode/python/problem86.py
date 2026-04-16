# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def partition(self, head: Optional[ListNode], x: int) -> Optional[ListNode]:
        l, r = ListNode(), ListNode()
        lans, rans = l, r
        while head is not None:
            if head.val < x:
                l.next = ListNode(head.val, None)
                l = l.next
            if head.val >= x:
                r.next = ListNode(head.val, None)
                r = r.next
            head = head.next
        l.next = rans.next
        return lans.next
