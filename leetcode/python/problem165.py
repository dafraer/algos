class Solution:
def compareVersion(self, version1: str, version2: str) -> int:
    r1 = version1.split(".")
    r2 = version2.split(".")
    for i in range(min(len(r1), len(r2))):
        if int(r1[i]) > int(r2[i]): return 1
        if int(r1[i]) < int(r2[i]): return -1
    if len(r1) > len(r2):
        for i in range(len(r2), len(r1)):
            if int(r1[i]) != 0: return 1
    elif len(r1) < len(r2):
        for i in range(len(r1), len(r2)):
            if int(r2[i]) != 0: return -1
    return 0
