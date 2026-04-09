func maxProduct(nums []int) int {
    n:=len(nums)
    ans:=nums[0]
    cur_max,cur_min:=1,1
    for i:=0;i<n;i++{
        prevmax,prevmin:=cur_max,cur_min
        cur_max=max(nums[i],max(prevmax*nums[i],prevmin*nums[i]))
        cur_min=min(nums[i],min(prevmax*nums[i],prevmin*nums[i]))
        ans=max(ans,cur_max)
    }
    return ans
}