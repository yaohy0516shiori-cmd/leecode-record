class Solution:
    def isValidSudoku(self, board: List[List[str]]) -> bool:
        seen=set()
        for i in range(9):
            for j in range(9):
                val=board[i][j]
                row=(val,'row',i)
                col=(val,'col',j)
                box=(val,'box',i//3,j//3)
                if val=='.':
                    continue
                if row in seen or col in seen or box in seen:
                    return False
                seen.add(row)
                seen.add(box)
                seen.add(col)
        return True
                    
                