 class Solution:
    def wordBreak(self, s: str, wordDict: List[str]) -> bool:
        m = {}
        def f(s):
            if s in wordDict: return True
            if s in m: return m[s]
            if len(s) <= 1: return False
            res = False
            for i in range(1,len(s)):
                if res: break
                res = f(s[0:i]) and f(s[i:])
            m[s] = res
            return res
        return f(s)
