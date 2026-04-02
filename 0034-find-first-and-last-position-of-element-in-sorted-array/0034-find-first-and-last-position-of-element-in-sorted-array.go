func searchRange(nums []int, target int) []int {
    left:=lowerbinary(nums,target)
    // if left side cannot find target, means there is no target in place
    if len(nums)==0 || nums[len(nums)-1]<target || nums[0]>target || nums[left]!=target{
        return []int{-1,-1}
    }
    
    right:=lowerbinary(nums,target+1)-1
    return []int{left,right}
}

// left find the smallest target place, right find the smallest target+1 place
func lowerbinary(nums []int, target int) int{
    l,r:=0,len(nums)
    for l<r{ // r does not participate in search, if r=len(nums)-1. l<=r
        mid:=(l+r)/2
        if nums[mid]<target{
            l=mid+1
        }else{
            r=mid
        }
    }
    return l
}