class Solution:
    def generateMatrix(self, n: int) -> List[List[int]]:
        visited = [[False for j in range(n)] for i in range(n)]
        matrix = [[0 for j in range(n)] for i in range(n)]
        cnt = 1

        def traverse(i, j, dir):
            nonlocal cnt
            matrix[i][j] = cnt
            cnt += 1
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
        return matrix
