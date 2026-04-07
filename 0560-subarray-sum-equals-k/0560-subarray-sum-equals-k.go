func subarraySum(nums []int, k int) int {
    res:=map[int]int{0:1}
    prefix:=0
    ans:=0
    for _, num := range nums{ // equal to i:=0;i<len(nums);i++
        prefix+=num
        ans+=res[prefix-k] //res[0]=1, or we cannot count prefix=k
        res[prefix]++
    }
    return ans
}