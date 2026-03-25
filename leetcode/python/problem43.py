class Solution:
    def multiply(self, num1: str, num2: str) -> str:
        if num1 == "0" or num2 == "0":
            return "0"
        res = []
        p1, p2 = len(num1) - 1, len(num2) - 1
        # multiply them
        while p2 >= 0:
            d = deque()
            transfer = 0
            while p1 >= 0:
                tmp = int(num1[p1]) * int(num2[p2]) + transfer
                transfer = tmp // 10
                d.appendleft(str(tmp % 10))
                p1 -= 1
            if transfer > 0:
                d.appendleft(str(transfer))
            res.append("".join(d))
            p1 = len(num1) - 1
            p2 -= 1
        maxlen = len(num1) + len(res)

        for i in range(len(res)):
            res[i] = res[i] + ("0" * i)
            if len(res[i]) < maxlen:
                res[i] = ("0" * (maxlen - len(res[i]))) + res[i]

        # sum them up
        ans = deque()
        p = -1
        transfer = 0
        while p >= -len(res[0]):
            tmp = transfer
            for x in res:
                tmp += int(x[p])
            p -= 1
            transfer = tmp // 10
            ans.appendleft(str(tmp % 10))
        if transfer > 0:
            ans.appendleft(str(transfer))
        realans = list(ans)
        for i in range(len(realans)):
            if realans[i] != "0":
                realans = realans[i:]
                break
        return "".join(realans)
