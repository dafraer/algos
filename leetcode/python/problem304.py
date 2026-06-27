class NumMatrix:

    def __init__(self, matrix: List[List[int]]):
        self.s = [list(itertools.accumulate(x)) for x in matrix]
        for i in range(1, len(self.s)):
            for j in range(len(self.s[i])):
                self.s[i][j] += self.s[i-1][j]

    def sumRegion(self, row1: int, col1: int, row2: int, col2: int) -> int:
        sm = self.s[row2][col2]
        if row1 > 0 and col1 > 0:
            return sm - self.s[row1-1][col2] - (self.s[row2][col1-1]-self.s[row1-1][col1-1])
        if row1 > 0: return sm - self.s[row1-1][col2]
        if col1 > 0: return sm - self.s[row2][col1-1]
        return sm
        

# Your NumMatrix object will be instantiated and called as such:
# obj = NumMatrix(matrix)
# param_1 = obj.sumRegion(row1,col1,row2,col2)