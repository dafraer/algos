# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class BSTIterator:
    def __init__(self, root: Optional[TreeNode]):
        self.root = root
        self.ptr = -math.inf

    def next(self) -> int:
        res = self.dfs(self.root, self.ptr)
        if res is not None:
            self.ptr = res
        return res

    def hasNext(self) -> bool:
        if self.dfs(self.root, self.ptr) is not None:
            return True
        return False

    def dfs(self, root, val) -> int:
        if not root:
            return None
        l = self.dfs(root.left, val)
        if l is not None:
            return l

        if root.val > val:
            return root.val

        return self.dfs(root.right, val)


# Your BSTIterator object will be instantiated and called as such:
# obj = BSTIterator(root)
# param_1 = obj.next()
# param_2 = obj.hasNext()
