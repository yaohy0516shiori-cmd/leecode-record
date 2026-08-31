class Solution:
    def solveNQueens(self, n: int) -> List[List[str]]:
        res=[ ['.' for _ in range (n)] for _ in range (n)]
        ans=[]

        def dfs(x):
            if x==n:
                ans.append([''.join(r) for r in res])
                return
            for i in range (n):
                if is_valid(i,x,res,n):
                    res[x][i]='Q'
                    dfs(x+1)
                    res[x][i]='.'
        
        def is_valid(col, row, res, n):
            for i in range(row):
                if res[i][col] == "Q":
                    return False

            i, j = row - 1, col + 1
            while i >= 0 and j < n:
                if res[i][j] == "Q":
                    return False
                i -= 1
                j += 1
            i,j=row-1,col-1
            while i >= 0 and j >= 0:
                if res[i][j]=='Q':
                    return False
                i-=1
                j-=1
            return True

        dfs(0)
        return ans