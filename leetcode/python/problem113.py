# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def pathSum(self, root: Optional[TreeNode], targetSum: int) -> List[List[int]]:
        res = []
        cur = []

        def traverse(node, cur_sum):
            if not node:
                return
            cur_sum += node.val
            cur.append(node.val)
            if cur_sum == targetSum and not node.left and not node.right:
                res.append(cur.copy())
            traverse(node.left, cur_sum)
            traverse(node.right, cur_sum)
            cur_sum -= cur[-1]
            cur.pop()

        traverse(root, 0)
        return res
