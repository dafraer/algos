class Solution:
    def maxProduct(self, nums: List[int]) -> int:
        res = max(nums)

        mn, mx = 1, 1

        for n in nums:
            tmp = mx * n
            mx = max(tmp, n * mn, n)
            mn = min(tmp, n * mn, n)
            res = max(res, mx)
        return res
