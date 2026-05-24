# Didn't solve
class Solution:
    def largestNumber(self, nums: List[int]) -> str:
        s = [str(x) for x in nums]

        def compare(a, b):
            if int(a + b) > int(b + a):
                return -1
            return 1

        s.sort(key=cmp_to_key(compare))
        if s[0] == "0":
            return "0"
        return "".join(s)
