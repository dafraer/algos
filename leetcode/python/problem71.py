class Solution:
    def simplifyPath(self, path: str) -> str:
        s = path.split("/")
        s = [x for x in s if len(x) > 0 and x != "."]
        i = 0
        while i < len(s):
            if s[i] == "..":
                s.pop(i)
                if i - 1 >= 0:
                    s.pop(i - 1)
                    i -= 1
                i -= 1
            i += 1
        return "/" + "/".join(s)
