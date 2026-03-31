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
            nums[start],nums[i]=nums[i],nums[start]
            dfs(start+1,end)
            nums[start],nums[i]=nums[i],nums[start]
        }
    }
    dfs(0,len(nums))
    return res
}