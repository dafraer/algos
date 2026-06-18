 class Solution:
    def hIndex(self, citations: List[int]) -> int:
        ans = 0
        for i in range(len(citations)):
            ans = max(ans, min(len(citations)-i, citations[i]))
        return ans
