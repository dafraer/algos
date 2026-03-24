class Solution:
    def countAndSay(self, n: int) -> str:
        def rle(s: str) -> str:
            compressed = []
            cnt = 1
            for i in range(0, len(s)):
                if (i < len(s) - 1) and (s[i] == s[i + 1]):
                    cnt += 1
                elif i < len(s) - 1:
                    compressed.append(str(cnt) + s[i])
                    cnt = 1
                elif cnt > 1:
                    compressed.append(str(cnt) + s[i])
                else:
                    compressed.append(str(1) + s[i])
            return "".join(compressed)

        # print(rle("3322251"))

        ans = ["1" for i in range(n + 1)]
        for i in range(2, n + 1):
            ans[i] = rle(ans[i - 1])
        return ans[n]
