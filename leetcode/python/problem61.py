# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def rotateRight(self, head: Optional[ListNode], k: int) -> Optional[ListNode]:
        # Edge cases
        if head == None:
            return None
        if k == 0:
            return head

        # Find length of the list
        h1 = head
        length = 0
        last = None
        while h1 is not None:
            length += 1
            last = h1
            h1 = h1.next

        k = k % length
        h2 = head
        last.next = head
        for i in range(1, length - k):
            h2 = h2.next

        ans = h2.next
        h2.next = None

        return ans
