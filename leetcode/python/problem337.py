# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def rob(self, root: Optional[TreeNode]) -> int:
        m = {}

        def f(node):
            if node in m:
                return m[node]
            if not node:
                return 0
            skip = f(node.left) + f(node.right)
            no_skip = node.val
            if node.left:
                no_skip += f(node.left.left) + f(node.left.right)
            if node.right:
                no_skip += f(node.right.left) + f(node.right.right)
            m[node] = max(skip, no_skip)
            return max(skip, no_skip)

        return f(root)
