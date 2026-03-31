func permute(nums []int) [][]int {
    res:=[][]int{}
    if len(nums)==1{
        return append(res,append([]int{},nums...))
    }
    var dfs func(start,end int)
    dfs=func(start,end int){
        if start==len(nums){
            res=append(res,append([]int{},nums...))
            return
        }
        for i:=start;i<end;i++{
            // fix the place before start and exchange, use loop to list all possibility
            nums[start],nums[i]=nums[i],nums[start]
            dfs(start+1,end)
            // backtrack
            nums[start],nums[i]=nums[i],nums[start]
        }
    }
    dfs(0,len(nums))
    return res
}