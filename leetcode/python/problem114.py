# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def flatten(self, root: Optional[TreeNode]) -> None:
        """
        Do not return anything, modify root in-place instead.
        """

        def build(node):
            if not node:
                return
            if not node.left:
                build(node.right)
                return
            if not node.right:
                build(node.left)
                node.right = node.left
                node.left = None
                return

            build(node.left)
            l = node.left
            r = node.right
            node.right = l
            while l.right:
                l = l.right
            l.right = r
            node.left = None
            build(r)

        build(root)
