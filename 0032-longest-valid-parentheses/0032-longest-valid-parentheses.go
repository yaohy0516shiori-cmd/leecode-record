func longestValidParentheses(s string) int {
    res:=[]int{-1}
    ans:=0
    for i:=0;i<len(s);i++{
        if s[i]=='('{
            res=append(res,i)
        }else{
            res=res[:len(res)-1]
            if len(res)==0{
                res=append(res,i)
            }else{
                ans=max(ans,i-res[len(res)-1])
            }
        }
    }
    return ans
}