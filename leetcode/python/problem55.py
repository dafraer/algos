 class Solution:
    def canJump(self, nums: List[int]) -> bool:
        mx = nums[0]
        i = 0
        while i <= mx and i < len(nums):
            mx = max(mx, i+nums[i])
            i+=1
        return mx >= len(nums)-1
            
