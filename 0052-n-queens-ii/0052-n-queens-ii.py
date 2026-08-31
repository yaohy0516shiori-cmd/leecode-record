class Solution:
    def totalNQueens(self, n: int) -> int:
        cols=set()
        right=set()
        left=set()
        count=0

        def dfs(row):
            nonlocal count
            if row==n:
                count+=1
                return
            
            for col in range(n):
                if col in cols:
                    continue
                if col+row in right :
                    continue
                if row-col in left:
                    continue
                
                cols.add(col)
                right.add(col+row)
                left.add(row-col)
                dfs(row+1)
                cols.remove(col)
                right.remove(col+row)
                left.remove(row-col)
        dfs(0)
        return count