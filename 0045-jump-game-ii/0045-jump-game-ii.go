func jump(nums []int) int {
    // case: nums := []int{2,3,4,1,1,1}
    ans:=0
    currentend:=0 // board of current place can reach, when reach, ans+1
    far:=0
    if len(nums)==1{
        return 0
    }
    for i:=0;i<len(nums)-1;i++{
        far=max(far,i+nums[i]) // farest place from this position
        if far==len(nums)-1{
            ans++
            return ans
        }
        if i==currentend{ // you must jump right now
            ans++
            currentend=far
        }
    }
    return ans
}