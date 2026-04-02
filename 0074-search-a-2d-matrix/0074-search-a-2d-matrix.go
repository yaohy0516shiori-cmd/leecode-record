func searchMatrix(matrix [][]int, target int) bool {
    if len(matrix[0])==0 || len(matrix)==0{
        return false
    }
    n:=len(matrix[0])
    lb,rb:=0,len(matrix)

    for lb<rb{
        mid_big:=(lb+rb)/2
        if matrix[mid_big][0]<=target && matrix[mid_big][n-1]>=target{
            return binary(matrix[mid_big],target)
        }else if matrix[mid_big][0]>target{
            rb=mid_big // shouldn't be mid_big-1, right place does not participate in search. 
            // if use mid_big-1, r initial should be len(nums)-1, l<=r (we need to search all place)
        }else if matrix[mid_big][n-1]<target{
            lb=mid_big+1
        }
    }
    return false
}

func binary(nums []int, target int) bool{
    ls,rs:=0,len(nums)
    for ls<rs{
        mid:=(ls+rs)/2
        if nums[mid]==target{
            return true
        }else if nums[mid]>target{
            rs=mid
        }else{
            ls=mid+1
        }
    }
    return false
}