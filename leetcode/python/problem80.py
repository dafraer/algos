class Solution:
    def removeDuplicates(self, nums: List[int]) -> int:
        nums.sort()
        i = 2
        while i < len(nums):
            if nums[i] == nums[i - 1] == nums[i - 2]:
                del nums[i]
                i -= 1
            i += 1
