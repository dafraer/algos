 class Solution:
    def nthUglyNumber(self, n: int) -> int:
        res = [1]
        added = set()
        for i in range(3*n):
            if res[i]*2 not in added:
                added.add(res[i]*2)
                res.append(res[i]*2)
            if res[i]*3 not in added:
                res.append(res[i]*3)
                added.add(res[i]*3)
            if res[i]*5 not in added:
                res.append(res[i]*5)
                added.add(res[i]*5)
        
        res.sort()
        return res[n-1]
