func longestPalindrome(s string) string {
    // dp[i,j]=dp[i+1,j-1] && s[i]==s[j]
    n:=len(s)
    ans_l,ans_r:=0,0

    // odd
    for i:=0;i<n;i++{
        l,r:=i,i
        for l>=0 && r<n && s[l]==s[r]{
            l--
            r++
            if r-l-1>ans_r-ans_l{
                ans_l,ans_r=l+1,r // left close, right open window
            }
        }
    }

    // even
    for i:=0;i<n;i++{
        l,r:=i,i+1
        for l>=0 && r<n && s[l]==s[r]{
            l--
            r++
            if r-l-1>ans_r-ans_l{
                ans_l,ans_r=l+1,r // left close, right open window
            }
        }
    }
    return s[ans_l:ans_r]
}