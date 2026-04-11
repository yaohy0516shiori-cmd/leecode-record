func canPartition(nums []int) bool {
    /*
    i: sum(nums)/2; sum must be even
    dp[i]: whether there is sum of subset equal to i
    dp[i]=dp[i-j]||dp[i] (choose j or not choose j, whether there is sum of subset equal to i)
    base case: dp[0]==true, when i==0, we choose nothing is fine
    it is a 0-1 package question(each element can be chosen for only 1 time)
    */
    ans:=0
    for i:=0;i<len(nums);i++{
        ans+=nums[i]
    }
    if ans%2==1{
        return false
    }
    ans=ans/2
    dp:=make([]bool,ans+1)
    dp[0]=true
    for _, num := range nums {
        for i := ans; i >= num; i-- {
            // 2D version: dp[k][i]=dp[k-1][i] || dp[k-1][i-num]; choose kth element or not
            dp[i] = dp[i] || dp[i-num]
        }
    }
    return dp[ans]
}