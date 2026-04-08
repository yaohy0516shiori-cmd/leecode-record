func fib(n int) int {
    if n==0{
        return 0
    }
    if n==1{
        return 1
    }
    // F(n) is start from zero, total n+1
    res:=make([]int,n+1)
    res[0]=0
    res[1]=1
    for i:=2;i<n+1;i++{
        res[i]=res[i-1]+res[i-2]
    }
    return res[n]
}