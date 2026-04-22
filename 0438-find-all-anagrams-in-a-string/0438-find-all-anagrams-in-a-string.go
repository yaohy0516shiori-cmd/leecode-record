func findAnagrams(s string, p string) []int {
    count:=make(map[byte]int)
    for i:=0;i<len(p);i++{
        count[p[i]]++
    }
    res:=[]int{}
    left:=0
    need:=len(p)
    if len(p)>len(s){
        return []int{}
    }
    // 
    for right:=0;right<len(s);right++{
        if count[s[right]]>0{
            need--
        }
        count[s[right]]--
        if right-left+1>len(p){
            if count[s[left]]>=0{
                need++
            }
            count[s[left]]++
            left++
        }
        if right-left+1==len(p) && need==0{
            res=append(res,left)
        }
    }
    return res
}