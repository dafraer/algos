class Solution:
    def findRepeatedDnaSequences(self, s: str) -> List[str]:
        if len(s) < 10:
            return []
        d = {}
        for i in range(len(s) - 9):
            d.setdefault(s[i : i + 10], 0)
            d[s[i : i + 10]] += 1

        res = []
        for k, v in d.items():
            if v > 1:
                res.append(k)
        return res
