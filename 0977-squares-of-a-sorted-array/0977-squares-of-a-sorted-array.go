func sortedSquares(nums []int) []int {
    n:=len(nums)
    left:=0
    right:=n-1
    ans:=make([]int,n)
    pos:=n-1
    for left<=right{
        ls:=nums[left]*nums[left]
        rs:=nums[right]*nums[right]
        if ls<rs{
            ans[pos]=rs
            right--
        }else{
            ans[pos]=ls
            left++
        }
        pos--
    }
    return ans
}