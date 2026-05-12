class Solution:
    def singleNumber(self, nums: List[int]) -> int:
        nums.sort()
        for i in range(0, len(nums) - 3, 3):
            a, b, c = nums[i], nums[i + 1], nums[i + 2]
            if a == b == c:
                continue
            if a == b:
                return c
            if b == c:
                return a
            if a == c:
                return b
        return nums[-1]
