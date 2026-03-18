class Solution:
    def threeSumClosest(self, nums: List[int], target: int) -> int:
        nums.sort()
        ansSum = nums[0] + nums[1] + nums[2]
        minDiff = abs(ansSum - target)
        for i, v in enumerate(nums):
            l, r = i + 1, len(nums) - 1
            while l < r:
                s = nums[l] + nums[r] + v
                diff = abs(s - target)
                if diff < minDiff:
                    ansSum = s
                    minDiff = diff
                if s > target:
                    r -= 1
                elif s < target:
                    l += 1
                else:
                    return s
        return ansSum
