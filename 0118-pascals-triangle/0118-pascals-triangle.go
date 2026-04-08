func generate(numRows int) [][]int {
    dp:=make([][]int,numRows)
    for i:=0;i<numRows;i++{
        dp[i]=make([]int,i+1)
        for j:=0;j<i+1;j++{
            dp[i][j]=1
        }
    }
    for i:=2;i<numRows;i++{
        for j:=1;j<i;j++{
            dp[i][j]=dp[i-1][j-1]+dp[i-1][j]
        }
    }
    return dp
}