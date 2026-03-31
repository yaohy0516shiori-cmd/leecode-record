func partition(s string) [][]string {
    res:=[][]string{}
    path:=[]string{}
    var dfs func(start int)
    dfs=func(start int){
        if start==len(s){
            res=append(res,append([]string{},path...))
            return
        }
        for i:=start;i<len(s);i++{
            sub:=s[start:i+1]  // just care what you cut down first, the rest will be deal in next step
            if !palindrome(sub){
                continue
            }
            path=append(path,sub)
            dfs(i+1)
            path=path[:len(path)-1]

        }
    }
    dfs(0)
    return res
}

func palindrome(s string) bool{
    lenth:=len(s)
    for i:=0;i<lenth/2;i++{
        if s[i]!=s[lenth-i-1]{
            return false
        }
    }
    return true
}