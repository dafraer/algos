# Didnt solve myself
class Solution:
    def maximumGap(self, nums: List[int]) -> int:
        if len(nums) < 2:
            return 0
        mx, mn = 0, math.inf
        for x in nums:
            if x > mx:
                mx = x
            if x < mn:
                mn = x
        b = max(1, math.ceil((mx - mn) / (len(nums) - 1)))
        buckets = [[-math.inf, math.inf] for _ in range((mx - mn) // b + 1)]
        for x in nums:
            buckets[(x - mn) // b][0] = max(buckets[(x - mn) // b][0], x)
            buckets[(x - mn) // b][1] = min(buckets[(x - mn) // b][1], x)
        ans_mx = 0
        prev_max = mn  # start from the global min
        for bmax, bmin in buckets:
            if bmax == -math.inf:
                continue  # empty bucket — skip but DON'T reset prev_max
            ans_mx = max(ans_mx, bmin - prev_max)
            prev_max = bmax  # only advance over non-empty buckets
        return ans_mx
