func longestConsecutive(nums []int) int {
    if len(nums)==0{
        return 0
    }
    hashmap:=make(map[int]bool)
    longest:=0
    for _,x :=range nums{
        hashmap[x]=true
    }
    for x:=range hashmap{
        if !hashmap[x-1]{
            cur:=x
            count:=1
            for hashmap[cur+1]{
                count++
                cur++
            }
            longest=max(longest,count)
        }
    }
    return longest
}