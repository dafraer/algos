class Solution:
    def uniquePathsWithObstacles(self, obstacleGrid: List[List[int]]) -> int:
        if obstacleGrid[0][0] == 1:
            return 0
        res = [[0 for x in y] for y in obstacleGrid]
        res[0][0] = 1
        for i in range(len(obstacleGrid)):
            for j in range(len(obstacleGrid[i])):
                if obstacleGrid[i][j] == 1:
                    continue
                if i > 0:
                    res[i][j] += res[i - 1][j]
                if j > 0:
                    res[i][j] += res[i][j - 1]
        return res[-1][-1]
