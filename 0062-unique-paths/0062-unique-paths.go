func uniquePaths(m int, n int) int {
    // dp[i][j]=dp[i][j-1]+dp[i-1][j]
    // only move downwards or rightwards
    // dp[0][j]=1, only 1 way
    // dp[i][0]=1, only 1 way
    dp:=make([][]int,m)
    for i:=0;i<m;i++{
        dp[i]=make([]int,n)
        for j:=0;j<n;j++{
            dp[0][j]=1
        }
        dp[i][0]=1
    }
    for i:=1;i<m;i++{
        for j:=1;j<n;j++{
            dp[i][j]=dp[i][j-1]+dp[i-1][j]
        }
    }
    return dp[m-1][n-1]
}
/*
func uniquePaths(m int, n int) int {
    dp := make([]int, n)

    // 第一行全是 1
    for j := 0; j < n; j++ {
        dp[j] = 1
    }

    // 从第二行开始
    for i := 1; i < m; i++ {
        for j := 1; j < n; j++ {
            dp[j] = dp[j] + dp[j-1]
        }
    }

    return dp[n-1]
}
*/