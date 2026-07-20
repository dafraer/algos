class Solution:
    def numDecodings(self, s: str) -> int:
        if s[0] == "0":
            return 0
        dp = [0 for _ in s]
        dp[0] = 1
        for i in range(1, len(s)):
            r = 0
            if 10 <= int(s[i - 1 : i + 1]) <= 26:
                if i > 1:
                    dp[i] += dp[i - 2]
                else:
                    dp[i] += 1
            if 1 <= int(s[i]) <= 9:
                dp[i] += dp[i - 1]
        return dp[-1]
