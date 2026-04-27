# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def levelOrderBottom(self, root: Optional[TreeNode]) -> List[List[int]]:
        q = deque()
        res = []

        def traverse(root, level):
            if not root:
                return
            if level >= len(res):
                res.append([root.val])
            else:
                res[level].append(root.val)
            q.append((root.left, level + 1))
            q.append((root.right, level + 1))

        q.append((root, 0))
        while len(q) > 0:
            node, lvl = q.popleft()
            traverse(node, lvl)
        res.reverse()
        return res
