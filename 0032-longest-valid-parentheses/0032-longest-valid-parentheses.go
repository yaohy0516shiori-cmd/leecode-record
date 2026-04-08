func longestValidParentheses(s string) int {
    // // use stack, if () at the beignning, ) locates in 1st position, ans calculate the length = 1-(-1)
    // res:=[]int{-1}
    // ans:=0
    // for i:=0;i<len(s);i++{
    //     // only store ( position, and pop up when meet )
    //     if s[i]=='('{
    //         res=append(res,i)
    //     }else{
    //         res=res[:len(res)-1]
    //         // if stack is nil, means )'s position is the new beignning
    //         if len(res)==0{
    //             res=append(res,i)
    //         }else{
    //             ans=max(ans,i-res[len(res)-1])
    //         }
    //     }
    // }
    // return ans
    
    
    
    n := len(s)
    if n == 0 {
        return 0
    }

    dp := make([]int, n)
    ans := 0
    // s[i]=(,dp[i]=0
    for i := 1; i < n; i++ {
        if s[i] == ')' {
            // s[i]=) && s[i-1]=(, dp[i]=dp[i-2]+2
            if s[i-1] == '(' {
                dp[i] = 2
                if i >= 2 {
                    dp[i] += dp[i-2]
                }
            } else {
                // s[i]=) && s[i-1]=), dp[i]=dp[i-dp[i-1]-1]+2+dp[i-1]
                j := i - dp[i-1] - 1
                if j >= 0 && s[j] == '(' {
                    dp[i] = dp[i-1] + 2
                    if j >= 1 {
                        dp[i] += dp[j-1]
                    }
                }
            }

            if dp[i] > ans {
                ans = dp[i]
            }
        }
    }

    return ans
}