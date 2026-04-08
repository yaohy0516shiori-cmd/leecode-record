func climbStairs(n int) int {
    // dp[i]=dp[i-1]+dp[i-2]
    // dp[i] is i number of methods, not length of path
    // eg: n=3, dp[3]=dp[2]+dp[1], dp[1]=1 dp[0]=1, dp[2]=dp[1]+dp[0]=2 (1+1,0+2)
    // it means if you are in the i-2 stage or i-1 stage, you can acheive the ist stage
    dp:=make([]int,n+1)
    dp[0]=1
    dp[1]=1
    for i:=2;i<n+1;i++{
        dp[i]=dp[i-1]+dp[i-2]
    }
    return dp[n]
}