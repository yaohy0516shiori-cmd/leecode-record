func exist(board [][]byte, word string) bool {
    lenth:=len(word)
    n:=len(board[0])
    m:=len(board)
    var dfs func(x,y,k int) bool
    dfs=func(x,y,k int) bool{
        if x>=m || y>=n  || x<0 || y<0 || board[x][y]!=word[k] {
            return false
        }
        if k==lenth-1{
            return true  // match
        }
        board[x][y]=0 // mark if it has been used, first sentence judge whether it is matched
        res:= dfs(x+1,y,k+1) || dfs(x-1,y,k+1) || dfs(x,y+1,k+1) || dfs(x,y-1,k+1)  // four direction
        board[x][y]=word[k] // backtrack
        return res // one true all true
    }
    for i:=0;i<m;i++ {
        for j:=0;j<n;j++{
            if dfs(i,j,0){
                return true
            }
        }
    }
    return false
}