#didnt solve
class NumArray:

    def __init__(self, nums: List[int]):
        n = len(nums)
        self.tree = [0]*n+nums
        for i in range(n-1, 0, -1):
            self.tree[i] = self.tree[2*i] + self.tree[(2*i+1)]

    def update(self, index: int, val: int) -> None:
        self.tree[len(self.tree)//2+index] = val
        ind = (len(self.tree)//2+index)//2
        while ind > 0:
            self.tree[ind] = self.tree[2*ind] + self.tree[(2*ind+1)]
            ind //= 2

    def sumRange(self, left: int, right: int) -> int:
        l = left + len(self.tree)//2
        r = right + len(self.tree)//2
        range_sum = 0
        
        while l <= r:
            if l % 2 == 1:
                    range_sum += self.tree[l]
                    l += 1
            if r % 2 == 0:
                    range_sum += self.tree[r]
                    r -= 1
            l //= 2
            r //= 2
            
        return range_sum
        

# Your NumArray object will be instantiated and called as such:
# obj = NumArray(nums)
# obj.update(index,val)
# param_2 = obj.sumRange(left,right)