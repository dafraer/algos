class Solution:
    def searchMatrix(self, matrix: List[List[int]], target: int) -> bool:
        for i in range(len(matrix)):
            l, r = 0, len(matrix[i])
            while l < r:
                m = (r + l) // 2
                if matrix[i][m] == target:
                    return True
                elif matrix[i][m] > target:
                    r = m
                else:
                    l = m + 1
        return False
