class Solution:
    def maxProduct(self, words: List[str]) -> int:
        mx = 0
        s1 = ""
        s2 = ""
        sts = [set(x) for x in words]
        for i in range(len(words)):
            for j in range(i+1, len(words)):
                if not (sts[i] & sts[j]) and len(words[i])*len(words[j]) > mx:
                    mx = len(words[i])*len(words[j])
                    s1 = words[i]
                    s2 = words[j]
        return mx
