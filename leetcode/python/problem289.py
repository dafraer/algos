class Solution:
    def gameOfLife(self, board: List[List[int]]) -> None:
        """
        Do not return anything, modify board in-place instead.
        """
        ans = [[0 for _ in board[0]] for _ in board]

        def n_cnt(i, j):
            return (
                (i > 0 and board[i-1][j])
                + (i < len(board)-1 and board[i+1][j])
                + (j > 0 and board[i][j-1])
                + (j < len(board[i])-1 and board[i][j+1])
                + ((i > 0 and j > 0 ) and board[i-1][j-1])
                + ((i > 0 and j < len(board[i])-1) and board[i-1][j+1])
                + ((j > 0 and i < len(board)-1) and board[i+1][j-1])
                + ((j < len(board[i])-1 and i < len(board)-1) and board[i+1][j+1])
            )
            
        for i in range(len(board)):
            for j in range(len(board[0])):
                c = board[i][j]
                ncnt = n_cnt(i, j)
                if c and (ncnt < 2 or ncnt > 3): ans[i][j] = 0
                elif not c and ncnt == 3: ans[i][j] = 1
                else: ans[i][j] = c
        for i in range(len(board)):
            for j in range(len(board[0])):
                board[i][j] = ans[i][j]
                
        