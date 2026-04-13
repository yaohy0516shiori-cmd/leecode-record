func canJump(nums []int) bool {
    //maxreach: farest place that n[i] can reach
    maxreach:=0
    for i:=0;i<len(nums);i++{
        if i>maxreach{
            return false
        }
        maxreach=max(maxreach,nums[i]+i)
    }
    return true
}