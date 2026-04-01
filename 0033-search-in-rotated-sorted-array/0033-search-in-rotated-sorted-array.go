func search(nums []int, target int) int {
    if len(nums)==0{
        return -1
    }
    n:=len(nums)
    l,r:=0,len(nums)-1
    for l<r{
        mid:=(l+r)/2
        if nums[mid]>nums[r]{
            l=mid+1
        }else{
            r=mid
        }
    
    }
    point:=l
    if nums[point]<=target && target<=nums[n-1]{
        return binary(nums,point,n-1,target)
    }
    return binary(nums,0,point-1,target)
}

func binary(nums []int, l int, r int, target int) int{
    for l<=r{
        mid:=(l+r)/2
        if nums[mid]==target{
            return mid
        }
        if nums[mid]<target{
            l=mid+1
        }else{
            r=mid-1
        }
    }
    return -1
}