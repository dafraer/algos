# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    # Did not solve =(
    def buildTree(self, inorder: List[int], postorder: List[int]) -> Optional[TreeNode]:
        m = {v: i for i, v in enumerate(inorder)}

        def build(l, r):
            if l > r:
                return None
            root = TreeNode(postorder.pop())

            ind = m[root.val]
            root.right = build(ind + 1, r)
            root.left = build(l, ind - 1)
            return root

        return build(0, len(inorder) - 1)
