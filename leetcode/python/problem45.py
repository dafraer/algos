class Solution:
    def jump(self, nums: List[int]) -> int:
        dists = [math.inf for x in nums]
        dists[0] = 0
        for i in range(0, len(nums)):
            for j in range(i, min(i + nums[i] + 1, len(nums))):
                dists[j] = min(dists[j], dists[i] + 1)
        return dists[-1]
