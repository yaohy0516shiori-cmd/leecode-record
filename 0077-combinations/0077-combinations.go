func combine(n int, k int) [][]int {
    res:=[][]int{}
    path:=[]int{}
    var dfs func(start int)
    dfs=func(start int){
        if len(path)==k{
            res=append(res,append([]int{},path...))
            return
        }
        for i:=start;i<n;i++{
            path=append(path,i+1)
            dfs(i+1)
            path=path[:len(path)-1]
        }
    }
    dfs(0)
    return res
}