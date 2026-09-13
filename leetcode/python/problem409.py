 class Solution:
    def longestPalindrome(self, s: str) -> int:
        m = Counter(s)
        ans = 0
        ad = 0
        for _, v in m.items():
            if v % 2: ad = 1
            ans += v - (v % 2)
        return ans+ad
            
