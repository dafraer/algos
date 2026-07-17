class Solution:
    def findTheDifference(self, s: str, t: str) -> str:
        m = Counter(s)
        m2 = Counter(t)
        for c in t:
            if m[c] != m2[c]:
                return c

