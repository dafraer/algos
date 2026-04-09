class Solution:
    def minDistance(self, word1: str, word2: str) -> int:
        grid = [[-1 for i in range(len(word2) + 1)] for j in range(len(word1) + 1)]

        def edit(i, j):
            if i >= len(word1) and j >= len(word2):
                return 0
            if grid[i][j] != -1:
                return grid[i][j]
            if i >= len(word1):
                grid[i][j] = edit(i, j + 1) + 1
                return grid[i][j]
            if j >= len(word2):
                grid[i][j] = edit(i + 1, j) + 1
                return grid[i][j]
            if word1[i] == word2[j]:
                grid[i][j] = edit(i + 1, j + 1)
                return grid[i][j]
            else:
                grid[i][j] = (
                    min(edit(i + 1, j + 1), min(edit(i + 1, j), edit(i, j + 1))) + 1
                )
                return grid[i][j]

        return edit(0, 0)
