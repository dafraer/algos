# Didnt solve myself fucking retard
class Solution:
    def solve(self, board: List[List[str]]) -> None:
        """
        Do not return anything, modify board in-place instead.
        """

        def traverse(i, j):
            if board[i][j] != "O":
                return
            board[i][j] = "#"
            if i < len(board) - 1:
                traverse(i + 1, j)
            if j < len(board[i]) - 1:
                traverse(i, j + 1)
            if i > 0:
                traverse(i - 1, j)
            if j > 0:
                traverse(i, j - 1)

        for i in range(0, len(board)):
            traverse(i, 0)
            traverse(i, len(board[i]) - 1)

        for j in range(0, len(board[0])):
            traverse(0, j)
            traverse(len(board) - 1, j)

        for i in range(len(board)):
            for j in range(len(board[i])):
                if board[i][j] == "#":
                    board[i][j] = "O"
                elif board[i][j] == "O":
                    board[i][j] = "X"
