class Solution:
    def reverse(self, x: int) -> int:
        print((2**31)-1)
        negative = False
        if x < 0:
            negative = True


        s = str(x)[::-1]
        if negative:
            s = s[:len(s)-1]
            if int(s) >= 2**31-1:
                return 0

            return -int(s)
        if int(s) >= 2**31-1:
                return 0

        return int(s)
    
