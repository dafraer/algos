class Solution:
    def countPrimes(self, n: int) -> int:
        if n < 3:
            return 0
        multiples = set()
        cnt = 0
        for i in range(2, n):
            if i in multiples:
                continue
            cnt += 1
            for j in range(i + i, n, i):
                multiples.add(j)
        return cnt
