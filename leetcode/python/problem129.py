# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def sumNumbers(self, root: Optional[TreeNode]) -> int:
        res = 0
        cur = []

        def dfs(root):
            if not root:
                return
            cur.append(root.val)
            if not root.left and not root.right:
                m = 1
                for x in reversed(cur):
                    nonlocal res
                    res += x * m
                    m *= 10
            dfs(root.left)
            dfs(root.right)
            cur.pop()

        dfs(root)
        return res
