class Solution:
    def solveSudoku(self, board: List[List[str]]) -> None:
        """
        Do not return anything, modify board in-place instead.
        """
        row=[set() for _ in range(9)]
        col=[set() for _ in range(9)]
        box=[[set() for _ in range(3)] for _ in range(3)] 
        emp=[]

        for i in range(9):
            for j in range(9):
                if board[i][j]!='.':
                    d=board[i][j]
                    row[i].add(d)
                    col[j].add(d)
                    box[i//3][j//3].add(d)
                else:
                    emp.append((i,j))
        
        def dfs(k):
            if k==len(emp):
                return True
            i,j=emp[k]
            for d in '123456789':
                if d in row[i] or d in col[j] or d in box[i//3][j//3]:
                    continue
                board[i][j]=d
                row[i].add(d)
                col[j].add(d)
                box[i//3][j//3].add(d)
                # core judgement
                if dfs(k+1):
                    return True

                board[i][j]='.'
                row[i].remove(d)
                col[j].remove(d)
                box[i//3][j//3].remove(d)

            return False

        dfs(0)
