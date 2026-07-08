# Didnt solve
class Solution:
    def wiggleSort(self, nums: List[int]) -> None:
        """
        Do not return anything, modify nums in-place instead.
        """
        if len(nums) == 1:
            return
        nums.sort()
        n = len(nums) // 2 + len(nums) % 2
        a = deque(nums[0:n])
        b = deque(nums[n:])
        for i in range(len(nums)):
            if i % 2:
                nums[i] = b.pop()
            else:
                nums[i] = a.pop()
