 class Solution:
    def combine(self, n: int, k: int) -> List[List[int]]:
        res = []
        def dp(curr, used, s, k):
            if k == 0:
                res.append(curr)
                return 
            for i in range(s, n+1):
                if used.get(i) is None:
                    curr_cp = curr.copy()
                    used_cp = used.copy()
                    curr_cp.append(i)
                    used_cp[i] = True
                    dp(curr_cp.copy(), used_cp.copy(), i+1, k-1)
        dp([], {}, 1, k)
        return res
