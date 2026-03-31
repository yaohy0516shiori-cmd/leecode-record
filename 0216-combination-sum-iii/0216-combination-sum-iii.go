func combinationSum3(k int, n int) [][]int {
    res:=[][]int{}
    path:=[]int{}
    var dfs func(start,remain int)
    dfs=func(start,remain int){
        if remain==0 && len(path)==k{
            res=append(res,append([]int{},path...))
            return
        }
        for i:=start;i<10;i++{
            if remain<i || len(path)>=k{
                break
            }
            path=append(path,i)
            dfs(i+1,remain-i)
            path=path[:len(path)-1]
        }
    }
    dfs(1,n)
    return res
}