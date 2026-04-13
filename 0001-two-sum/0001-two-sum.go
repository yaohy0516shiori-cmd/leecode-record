func twoSum(nums []int, target int) []int {
    hashtable:= make(map[int]int)
    for i,j := range nums{
        component:=target-j
        if idx,ok:=hashtable[component];ok{
            return []int{idx,i}
        }
        hashtable[j]=i
    }
    return nil
}