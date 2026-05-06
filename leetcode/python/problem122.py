 class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        mn_price = prices[0]
        mx_profit = 0
        for i in range(1, len(prices)):
            if prices[i]-mn_price > 0: 
                mx_profit += prices[i]-mn_price
                mn_price = prices[i]
            
            mn_price = min(mn_price, prices[i])
        return mx_profit
