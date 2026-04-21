func lengthOfLongestSubstring(s string) int {
    res:=make(map[byte]int)
    n:=len(s)
    left,right:=0,0
    ans:=0
    for right<n{
        res[s[right]]++
        for res[s[right]]>1{
            res[s[left]]--
            left++
        }
        ans=max(ans,right-left+1)
        right++
    }
    return ans
}