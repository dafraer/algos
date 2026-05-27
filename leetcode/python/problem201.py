class Solution:
    def rangeBitwiseAnd(self, left: int, right: int) -> int:
        i = 1 << 32
        start = False
        res = 0
        while i > 0:
            if (i | right) != right and (i | right) != left and not start:
                i >>= 1
                continue
            start = True
            if ((i | right) == right) != ((i | left) == left):
                break
            if ((i | right) == right) and ((i | left) == left):
                res |= i
            i >>= 1
        return res
