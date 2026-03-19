# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def swapPairs(self, head: Optional[ListNode]) -> Optional[ListNode]:
        if head == None or head.next == None:
            return head
        ans = head.next

        h1 = head
        h2 = head.next
        h3 = h2.next
        h2.next = h1
        h1.next = self.swapPairs(h3)

        return ans
