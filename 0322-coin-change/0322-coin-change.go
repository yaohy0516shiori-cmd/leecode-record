func coinChange(coins []int, amount int) int {
    /*
    dp[i]: the fewest number of coins make up amount
    dp[i]=dp[i-j]+1; j in coins
    */
    dp:=make([]int, amount+1)
    if amount==0{
        return 0
    }
    dp[0]=0
    for i:=1;i<amount+1;i++{
        dp[i]=amount+1
        for j:=0;j<len(coins);j++{
            if i>=coins[j]{
                dp[i]=min(dp[i],dp[i-coins[j]]+1)
            }
        }
    }
    if dp[amount]==amount+1{
        return -1
    }
    return dp[amount]
}