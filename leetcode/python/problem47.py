 class Solution:
    def permuteUnique(self, nums: List[int]) -> List[List[int]]:
        res = []
        perm = []
        cnt = {n:0 for n in nums}
        for x in nums:
            cnt[x] += 1
        
        def dfs():
            if len(nums) == len(perm):
                res.append(perm.copy())
                return
            for x in cnt:
                if cnt[x] > 0:
                    perm.append(x)
                    cnt[x] -= 1
                    dfs()
                    
                    cnt[x] += 1
                    perm.pop()
        dfs()
        return res
