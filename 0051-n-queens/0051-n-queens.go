func solveNQueens(n int) [][]string {
    res:=[][]string{}
    board:=make([][]byte,n) 
    // create 2-D slice? shape: n row(already has position but empty), but row is nil?
    for i:=0; i<n; i++{
        board[i]=make([]byte,n) 
    }
    for i:=0;i<n;i++{
        for j:=0;j<n;j++{
            board[i][j]='.' 
            // what is the difference between "" and ''? '' is rune type."" is string
        }
    }

    var dfs func(rows int) 
    dfs=func(rows int){
        if rows==n{
            temp:=make([]string,n)
            for i:=0;i<n;i++{
                temp[i]=string(board[i])
            }
            res=append(res,temp) // doesn't need to copy
            return
        }
        for i:=0;i<n;i++{
            if is_valid(i,rows,board,n){
                board[rows][i]='Q'
                dfs(rows+1)
                board[rows][i]='.'
            }
        }

    }
    dfs(0)
    return res
}


func is_valid(col int, row int, board [][]byte, n int) bool{
    // col means current column, row means current row, n is slice columns

    // column check
    for i:=0; i<row; i++{  // not include current row, coz start from zero
        if board[i][col]=='Q'{
            return false
        }
    }
    // 45 degree check
    for i,j:=row-1,col-1;i>=0 && j>=0;i,j=i-1,j-1{
        if board[i][j]=='Q'{
            return false
        }
    }
    // 135 degree check
    for i,j:=row-1,col+1;i>=0 && j<n;i,j=i-1,j+1{
        if board[i][j]=='Q'{
            return false
        }
    }
    return true
}