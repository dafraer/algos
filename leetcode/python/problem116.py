"""
# Definition for a Node.
class Node:
    def __init__(self, val: int = 0, left: 'Node' = None, right: 'Node' = None, next: 'Node' = None):
        self.val = val
        self.left = left
        self.right = right
        self.next = next
"""


class Solution:
    def connect(self, root: "Optional[Node]") -> "Optional[Node]":
        q = deque()
        q.append(root)
        size = 0
        cnt = 1
        next = None
        while len(q) > 0:
            node = q.popleft()
            if len(q) > 0:
                next = q[0]
            if cnt == (1 << size):
                size += 1
                next = None
                cnt = 1
            else:
                cnt += 1
            if node:
                node.next = next
                q.append(node.left)
                q.append(node.right)
        return root
