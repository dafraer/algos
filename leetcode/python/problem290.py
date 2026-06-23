class Solution:
    def wordPattern(self, pattern: str, s: str) -> bool:
        words = s.split(' ')
        if len(pattern) != len(words): return False
        m1 = dict()
        m2 = dict()
        for i in range(len(pattern)):
            if pattern[i] in m1 and m1[pattern[i]] != words[i]: return False
            if words[i] in m2 and m2[words[i]] != pattern[i]: return False
            m1[pattern[i]] = words[i]
            m2[words[i]] = pattern[i]
        return True