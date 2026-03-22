class Solution:
    def nextPermutation(self, nums: List[int]) -> None:
        """
        Do not return anything, modify nums in-place instead.
        """
        first_decreasing = -1
        for i in reversed(range(len(nums))):
            if nums[i] > nums[i - 1]:
                first_decreasing = i - 1
                break
        if first_decreasing == -1:
            nums.reverse()
            return
        print(f"{first_decreasing=}")
        min_higher_ind = 0
        min_higher = math.inf
        for i in range(first_decreasing + 1, len(nums)):
            if nums[i] > nums[first_decreasing] and nums[i] < min_higher:
                min_higher = nums[i]
                min_higher_ind = i
                print(f"{min_higher=}, {min_higher_ind=}")
        nums[first_decreasing], nums[min_higher_ind] = (
            nums[min_higher_ind],
            nums[first_decreasing],
        )
        r = len(nums)
        l = first_decreasing + 1
        nums[l:r] = sorted(nums[l:r])
