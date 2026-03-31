class Solution:
    def spiralOrder(self, matrix: List[List[int]]) -> List[int]:
        visited = [[False for j in range(len(matrix[i]))] for i in range(len(matrix))]
        res = []

        def traverse(i, j, dir):
            res.append(matrix[i][j])
            visited[i][j] = True
            match dir:
                case "right":
                    if j + 1 >= len(matrix[i]) or visited[i][j + 1]:
                        if i + 1 >= len(matrix) or visited[i + 1][j]:
                            return
                        traverse(i + 1, j, "down")
                    else:
                        traverse(i, j + 1, "right")
                case "down":
                    if i + 1 >= len(matrix) or visited[i + 1][j]:
                        if j - 1 < 0 or visited[i][j - 1]:
                            return
                        traverse(i, j - 1, "left")
                    else:
                        traverse(i + 1, j, "down")
                case "left":
                    if j - 1 < 0 or visited[i][j - 1]:
                        if i - 1 < 0 or visited[i - 1][j]:
                            return
                        traverse(i - 1, j, "up")
                    else:
                        traverse(i, j - 1, "left")
                case "up":
                    if i - 1 < 0 or visited[i - 1][j]:
                        if j + 1 >= len(matrix) or visited[i][j + 1]:
                            return
                        traverse(i, j + 1, "right")
                    else:
                        traverse(i - 1, j, "up")

        traverse(0, 0, "right")
        return res
