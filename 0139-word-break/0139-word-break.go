func wordBreak(s string, wordDict []string) bool {
    //dp[i]=dp[j] && is exist s[j:i] in worddict, j<i; dp[0]=true(any dict can represent '')
    /*    s[j:i]
    i=0   "" 
    i=1   "c"
    i=2   "ca"
    i=3   "cat"
    i=4   "cats"
    i=5   "catsa"
    i=6   "catsan"
    i=7   "catsand"
    i=8   "catsando"
    i=9   "catsandog"
    i      0    1    2    3    4    5      6      7       8        9
    prefix ""   c    ca   cat  cats catsa catsan catsand catsando catsandog
    dp     T    F    F    T    T    F      F       T      F        F
    */
    // or dp[i]=dp[i+len(word)] && s[i:i+len(word)]==word, word in dict; dp[n]=true (back to front)
    dp:=make([]bool,len(s)+1)
    dp[0]=true
    wordset:=map[string]bool{}
    for _,w := range wordDict{
        wordset[w]=true
    }
    for i:=1;i<len(s)+1;i++{
        for j:=0;j<i;j++{
            if wordset[s[j:i]] && dp[j]{
                dp[i]=true
                break
            }else{
                dp[i]=false
            }
        }
    }
    return dp[len(s)]
}