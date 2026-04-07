class Solution:
    def minPathSum(self, grid: List[List[int]]) -> int:
        res = [[math.inf for cell in row] for row in grid]
        res[0][0] = 0
        for i in range(len(grid)):
            for j in range(len(grid[i])):
                if i > 0:
                    res[i][j] = min(res[i - 1][j], res[i][j])
                if j > 0:
                    res[i][j] = min(res[i][j - 1], res[i][j])
                res[i][j] += grid[i][j]
        return res[-1][-1]
