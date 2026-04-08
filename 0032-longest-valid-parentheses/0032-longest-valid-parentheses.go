func longestValidParentheses(s string) int {
    // use stack, if () at the beignning, ) locates in 1st position, ans calculate the length = 1-(-1)
    res:=[]int{-1}
    ans:=0
    for i:=0;i<len(s);i++{
        // only store ( position, and pop up when meet )
        if s[i]=='('{
            res=append(res,i)
        }else{
            res=res[:len(res)-1]
            // if stack is nil, means )'s position is the new beignning
            if len(res)==0{
                res=append(res,i)
            }else{
                ans=max(ans,i-res[len(res)-1])
            }
        }
    }
    return ans
}