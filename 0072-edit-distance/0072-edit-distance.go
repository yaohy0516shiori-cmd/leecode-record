/* 
    dp[i][j] represents the minimum number of operations required to 
    transform the first i characters of word1 into the first j characters of word2.
    base case: dp[i][0]=i; dp[0][j]=j
    explain: 
        ' ' r  o  s
    ' '  0  1  2  3   (r->' '; ro->' '; ros-> ' ') steps of word change to nil
    h    1  1  2  3
    o    2  2  1  2
    r    3  2  2  2
    s    4  3  3  2
    e    5  4  4  3
    operation: 1.repalce letter in A == replace letter in B  2.insert letter in A == delete letter in B  3.delete letter in A == insert letter in B
    dp[i][j]=min(dp[i-1][j],dp[i][j-1],dp[i-1][j-1])+1 || dp[i][j]=dp[i-1][j-1]
    */
func minDistance(word1 string, word2 string) int {
    dp:=make([][]int,len(word1)+1)
    for i:=0;i<len(word1)+1;i++{
        dp[i]=make([]int,len(word2)+1)
        dp[i][0]=i
    }
    for j:=1;j<len(word2)+1;j++{
        dp[0][j]=j
    }
    for i:=1;i<len(word1)+1;i++{
        for j:=1;j<len(word2)+1;j++{
            if word1[i-1]!= word2[j-1]{
                dp[i][j]=min(dp[i-1][j],dp[i][j-1],dp[i-1][j-1])+1
            }else{
                dp[i][j]=dp[i-1][j-1]  // word1[i]=word2[j]-->doesn't operate
            }
            
        }
    }
    return dp[len(word1)][len(word2)]
}