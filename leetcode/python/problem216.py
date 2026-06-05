class Solution:
    def combinationSum3(self, k: int, n: int) -> List[List[int]]:
        res = []
        curr = []

        def backtrack(sm, val):
            sm += val
            curr.append(val)

            if sm == n and len(curr) == k:
                res.append(curr.copy())
            if sm >= n or len(curr) >= k:
                sm -= val
                curr.pop()
                return

            for i in range(val + 1, 10):
                backtrack(sm, i)
            sm -= val
            curr.pop()

        for i in range(1, 10):
            backtrack(0, i)
        return res
