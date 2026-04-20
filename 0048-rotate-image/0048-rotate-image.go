func rotate(matrix [][]int)  {
    n:=len(matrix)-1
    mid:=n/2
    for i:=0; i<=mid; i++{
        for j:=0; j<=n; j++{
            temp:=matrix[i][j]
            matrix[i][j]=matrix[n-i][j]
            matrix[n-i][j]=temp
        }
    }
    for i:=0; i<=n; i++{
        for j:=i; j<=n; j++{
            temp:=matrix[i][j]
            matrix[i][j]=matrix[j][i]
            matrix[j][i]=temp
        }
    }
}