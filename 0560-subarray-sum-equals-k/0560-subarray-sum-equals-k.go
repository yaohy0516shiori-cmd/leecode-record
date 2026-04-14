func subarraySum(nums []int, k int) int {
    /*
    [1,2,3,4,5,5,7] k=10
    prefix[1,3,6,10,15,20,28]
    if prefix-k in hashmap, means there is an subarray sum==k
    if not, there is not subarray
    when we need subarray, prefix is a very good choice
    */
    res:=map[int]int{0:1}
    prefix:=0
    ans:=0
    for _, i := range nums{
        prefix+=i
        ans+=res[prefix-k]
        res[prefix]++
    }
    return ans 
}