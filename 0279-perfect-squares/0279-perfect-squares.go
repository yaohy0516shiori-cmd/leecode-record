func numSquares(n int) int {
    /*
    dp[i]=the least number of perfect squares that sum to i
    1 2 3 4 5 6 7 8 9 10 11 12 13
    1 2 3 1 2 3 4 2 1 2   3  3  2
    dp[i]=min(dp[i-j*j]+1); j*j<i
    */
    dp:=make([]int,n+1)
    dp[0]=0
    dp[1]=1
    if n==1{
        return 1
    }
    for i:=2;i<n+1;i++{
        mins:=10000
        for j:=1;j*j<=i;j++{
            temp:=dp[i-j*j]+1
            if temp<mins{
                mins=temp
            }
        }
        dp[i]=mins
    }
    return dp[n]
}