class Solution:
    def lengthOfLIS(self, nums: List[int]) -> int:
        m = {}

        def f(i):
            if i in m:
                return m[i]
            if i >= len(nums):
                return 0
            mx = 0
            for j in range(i, len(nums)):
                if nums[i] < nums[j]:
                    mx = max(mx, f(j))
            m[i] = mx + 1
            return mx + 1

        mx = 0
        for i in range(len(nums)):
            mx = max(mx, f(i))
        return mx
