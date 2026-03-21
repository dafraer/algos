class Solution:
    def divide(self, dividend: int, divisor: int) -> int:
        cnt = 0
        abs_divisor = abs(divisor)
        abs_dividend = abs(dividend)
        if abs_divisor == 1:
            return self.adjust_sign(dividend, divisor, dividend)

        d = dict()
        curr_divisor = abs_divisor
        while abs_dividend >= curr_divisor:
            for k, v in reversed(sorted(d.items())):
                if curr_divisor + k <= abs_dividend:
                    curr_divisor += k
                    cnt += v
                    break
            else:
                curr_divisor += abs_divisor
                cnt += 1
            d[curr_divisor] = cnt + 1
        return self.adjust_sign(dividend, divisor, cnt)

    def adjust_sign(self, dividend: int, divisor: int, ret_val: int):
        def adjust_val(v: int):
            if v > (2**31) - 1:
                return (2**31) - 1
            if v < -(2**31):
                return -(2**31)
            return v

        v = 0
        if (dividend < 0) ^ (divisor < 0):
            v = -abs(ret_val)
        else:
            v = abs(ret_val)
        return adjust_val(v)
