func searchMatrix(matrix [][]int, target int) bool {
    m,n:=len(matrix),len(matrix[0])
    if m==0 || n==0{
        return false
    }
    i,j:=0,n-1
    for i<m && j>=0{
        if matrix[i][j]==target{
            return true
        }else if matrix[i][j]>target{
            j--
        }else{
            i++
        }
    }
    return false
}