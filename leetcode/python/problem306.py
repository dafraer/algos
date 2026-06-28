class Solution:
    def isAdditiveNumber(self, num: str) -> bool:
        def backtrack(f, s, ind):
            if ind == len(num): return True
            ans = False
            for i in range(ind, len(num)):
                if (num[ind] != '0' or ind == i) and f+s == int(num[ind:i+1]):
                    ans = ans or backtrack(s, int(num[ind:i+1]), i+1)
            return ans

        for i in range(1, len(num)):
            for j in range(i+1, len(num)):
                if (num[i] != '0' or i == j-1) and (num[0] != '0' or i-1 == 0) and backtrack(int(num[0:i]), int(num[i:j]), j): return True

        return False