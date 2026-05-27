func convert(s string, numRows int) string {
    if numRows==1 || numRows>=len(s){
        return s
    }
    rows:=make([][]byte,numRows)
    cur:=0
    dir:=1
    for i:=0;i<len(s);i++{
        rows[cur]=append(rows[cur],s[i])
        if cur==0{
            dir=1
        }else if cur==numRows-1{
            dir=-1
        }
        cur+=dir
    }
    res:=make([]byte,0,len(s))
    for i:=0;i<numRows;i++{
        res=append(res,rows[i]...)
    }
    return string(res)
}