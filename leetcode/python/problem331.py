# Didnt solve
class Solution:
    def isValidSerialization(self, preorder: str) -> bool:
        vals = preorder.split(",")
        t = 1
        for v in vals:
            if not t:
                return False
            if v == "#":
                t -= 1
            else:
                t += 1
        return t == 0
