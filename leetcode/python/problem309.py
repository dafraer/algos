#Didnt solve
class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        dp = {}

        def f(i, b):
            if i >= len(prices): return 0
            if (i, b) in dp: return dp[(i, b)]
            if b: 
                buy = f(i+1, not b) - prices[i]
                cd = f(i+1, b)
                dp[(i, b)] = max(buy, cd)
            else: 
                sell = f(i+2, not b) + prices[i]
                cd = f(i+1, b)
                dp[(i, b)] = max(sell, cd)
            return dp[(i, b)]
        return f(0, True)