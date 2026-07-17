class Solution:
    def findTheDifference(self, s: str, t: str) -> str:
        m = set(s)
        for c in t:
            if c not in m:
                return c
