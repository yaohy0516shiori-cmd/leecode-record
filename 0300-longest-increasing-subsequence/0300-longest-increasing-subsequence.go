func lengthOfLIS(nums []int) int {
    /*
    dp[i]: the strictly increasing subsequence that ends at nums[i] because nums[i] cannot only compare with nums[i-1]
    10 9 2 5 3 7 101 18 
    1  1  1 2 2 3 4   4  
    for j<i, if nums[j]<nums[i], nums[i] can extend the subsequence at j, dp[i]=max(dp[i],dp[j]+1)
    base case: dp[i]=1 for everyone
    */
    n:=len(nums)
    if n==1{
        return 1
    }
    dp:=make([]int,n)
    ans:=1
    for i:=0;i<n;i++{
        dp[i]=1
        for j:=0;j<i;j++{
            if nums[j]<nums[i]{
                dp[i]=max(dp[i],dp[j]+1)
            }
        }
        ans=max(ans,dp[i]) //ans does not have to be in last position
    }
    return ans
}