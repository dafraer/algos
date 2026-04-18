# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def reverseBetween(
        self, head: Optional[ListNode], left: int, right: int
    ) -> Optional[ListNode]:
        h = head
        mid_ind = (right + left) // 2 + (1 - (right - left + 1) % 2)
        i = 1
        while i < mid_ind:
            if i >= left:
                dist = (mid_ind - i) * 2 - (1 - (right - left + 1) % 2)
                tmp = h
                for x in range(dist):
                    tmp = tmp.next
                tmp.val, h.val = h.val, tmp.val
            i += 1
            h = h.next

        return head
