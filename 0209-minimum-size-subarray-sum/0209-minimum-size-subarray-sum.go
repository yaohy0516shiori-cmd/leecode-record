func minSubArrayLen(target int, nums []int) int {
    left:=0
    ans:=0
    res:=len(nums)+1
    for right:=0;right<len(nums);right++{
        ans+=nums[right]
        for ans>=target{
            length:=right-left+1
            res=min(res,length)
            ans-=nums[left]
            left++
        }

    }
    if res==len(nums)+1{
        return 0
    }
    return res
}