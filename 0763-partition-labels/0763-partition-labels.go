func partitionLabels(s string) []int {
    hash:=make(map[byte]int)
    for i:=0;i<len(s);i++{
        hash[s[i]]=i
    }
    result:=[]int{}
    start,end:=0,0
    for i:=0;i<len(s);i++{
        end=max(end,hash[s[i]]) // the farest place of the char that you have already scanned
        if i==end{  
            // when the vector walk to the farest place of the char that you have already scanned, mean there should not be any same char behind
            result=append(result,end-start+1)
            start=i+1
        }
    }
    return result
}