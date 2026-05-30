class Solution:
    def minSubArrayLen(self, target: int, nums: List[int]) -> int:
        l = r = sm = 0
        mn = len(nums) + 1
        while r < len(nums) + 1:
            print(f"{l=} {r=} {sm=} {mn=}")
            if sm < target and r < len(nums):
                sm += nums[r]
                r += 1
            elif sm >= target and l < len(nums):
                mn = min(mn, (r - l))
                sm -= nums[l]
                l += 1
            else:
                break
        if mn != len(nums) + 1:
            return mn
        return 0
