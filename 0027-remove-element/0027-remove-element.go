func removeElement(nums []int, val int) int {
    left:=0
    right:=len(nums)-1
    for left<=right{
        for left<=right && nums[right]==val{
            right--
        }
        if left>right{
            break
        }
        if nums[left]==val{
            nums[left]=nums[right]
            right--
            left++
        }else{
            left++
        }
    }
    return right+1
}