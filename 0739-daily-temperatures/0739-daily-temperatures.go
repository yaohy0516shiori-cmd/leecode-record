func dailyTemperatures(temperatures []int) []int {
    stack:=[]int{}
    res:=make([]int,len(temperatures))
    for index,value:= range temperatures{
        for len(stack)!=0 && temperatures[stack[len(stack)-1]]<value{
            res[stack[len(stack)-1]]=index-stack[len(stack)-1]
            stack=stack[:len(stack)-1]
        }
        stack=append(stack,index)
    }
    return res
}