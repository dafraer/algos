class Solution:
    def trailingZeroes(self, n: int) -> int:
        sys.set_int_max_str_digits(99999)
        res = str(math.factorial(n))
        print(res)
        cnt = 0
        for i in range(len(res) - 1, -1, -1):
            if res[i] != "0":
                break
            cnt += 1
        return cnt
