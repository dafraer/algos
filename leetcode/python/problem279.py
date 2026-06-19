#didnt solve
class Solution:
    def numSquares(self, n: int) -> int:
        m = dict()
        def f(n):
            if n <= 0: return 0
            cnt = math.inf
            for i in range(1, n+1):
                if i*i > n: break
                if n-i*i in m: 
                    tmp = m[n-i*i]
                else: 
                    tmp = f(n-i*i)
                    m[n-i*i] = tmp
                cnt = min(cnt, tmp)
            return cnt+1
        return f(n)
