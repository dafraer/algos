class Solution:
    def searchRange(self, nums: List[int], target: int) -> List[int]:
        l, r = 0, len(nums)
        ans = [-1, -1]
        while l < r:
            m = (r + l) // 2
            if nums[m] > target:
                r = m
            elif nums[m] < target:
                l = m + 1
            else:
                l = r = m
                while (l >= 0 and nums[l] == target) or (
                    r < len(nums) and nums[r] == target
                ):
                    if l >= 0 and nums[l] == target:
                        ans[0] = l
                        l -= 1
                    if r < len(nums) and nums[r] == target:
                        ans[1] = r
                        r += 1
                break
        return ans
