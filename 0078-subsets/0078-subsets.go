func subsets(nums []int) [][]int {
    res:=[][]int{}
    path:=[]int{}
    var dfs func(index int)
    dfs=func(index int){
        res=append(res,append([]int{},path...))
        for i:=index;i<len(nums);i++{
            path=append(path,nums[i])
            dfs(i+1)  // not index+1 coz index doesn't change with loop
            path=path[:len(path)-1]
        }
    }
    dfs(0)
    return res
}