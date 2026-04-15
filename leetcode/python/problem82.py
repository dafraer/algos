# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def deleteDuplicates(self, head: Optional[ListNode]) -> Optional[ListNode]:
        h = head
        prev = ListNode()
        prev.next = h
        ans = prev
        while h is not None:
            val = h.val
            cur = h.next
            while cur is not None and cur.val == val:
                cur = cur.next
            if cur != h.next:
                prev.next = cur
                h = cur
            else:
                prev = h
                h = h.next
        return ans.next
