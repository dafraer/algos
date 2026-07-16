# Didnt solve
class Solution:
    def integerBreak(self, n: int) -> int:
        m = {}

        def dfs(a):
            if a == 1:
                return 1
            if a in m:
                return m[a]
            if a != n:
                mx = a
            else:
                mx = 0
            for i in range(1, a):
                mx = max(dfs(a - i) * dfs(i), mx)
            m[a] = mx
            return mx

        return dfs(n)
