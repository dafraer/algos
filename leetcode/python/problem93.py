 class Solution:
    def restoreIpAddresses(self, s: str) -> List[str]:
        res = []
        def b(ind, cur, cnt):
            if cnt == 3 and len(cur) == len(s)+3:
                res.append(cur)
                return
            if ind >= len(s) or cnt == 3:
                return
            lim = min(4, (len(s)-ind)+1)
            for i in range(ind+1, ind+lim):
                tmp = s[ind:i]
                a = int(tmp)
                if a >= 0 and a < 256 and not (tmp[0] == '0' and len(tmp) > 1):
                    if cur == '':
                        b(i, s[ind:i], cnt)
                    else:
                        b(i, ".".join([cur, s[ind:i]]), cnt+1)
        b(0, '', 0) 
        return res    
