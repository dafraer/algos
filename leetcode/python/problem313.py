#Didn't solve
class Solution:
    def nthSuperUglyNumber(self, n: int, primes: List[int]) -> int:
        heap = [1]
        added = set()
        for i in range(n):
            num = heapq.heappop(heap)
            if i == n-1: return num
            
            for x in primes:
                if x*num not in added:
                    added.add(x*num)
                    heapq.heappush(heap, x*num)