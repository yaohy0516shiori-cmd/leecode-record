import "sort"
func singleNumber(nums []int) int {
    sort.Ints(nums)
    for i:=0;i<(len(nums)-1);{
        if nums[i]==nums[i+1]{
            i+=2
        }else{
            return nums[i]
        }
    }
    return nums[len(nums)-1]
}