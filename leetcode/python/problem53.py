 class Solution:
    def maxSubArray(self, nums: List[int]) -> int:
        mx = dp_sum = nums[0]
        for i in range(1, len(nums)):
            dp_sum = max(dp_sum+nums[i], nums[i])
            mx = max(mx, dp_sum)
        return mx 
