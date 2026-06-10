class Solution:
    def majorityElement(self, nums: List[int]) -> List[int]:
        m = Counter(nums)
        res = [k for k, v in m.items() if v > len(nums) // 3]
        return res
