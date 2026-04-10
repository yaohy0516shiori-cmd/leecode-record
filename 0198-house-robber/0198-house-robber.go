func rob(nums []int) int {
    /*
    when you query dp[i], answer is the maximum money you can get from houses [0...i]
    if you steal i, dp[i]=dp[i-2]+s[i]
    if not, dp[i]=dp[i-1]
    dp[i]=max(dp[i-1],dp[i-2]+s[i])
    base case: 
    dp[0]=s[0]
    dp[1]=s[1]
    */
    n:=len(nums)
    dp:=make([]int,n)
    dp[0]=nums[0]
    if n==1{
        return dp[0]
    }
    dp[1]=max(nums[0],nums[1])
    for i:=2;i<n;i++{
        dp[i]=max(dp[i-1],dp[i-2]+nums[i])
    }
    return dp[n-1]
}