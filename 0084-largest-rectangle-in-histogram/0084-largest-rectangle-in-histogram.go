func largestRectangleArea(heights []int) int {
    // farest place can reach to left/right, find the first col that lower than previous
    // use stack store index, can query height, if current col < previous, right place
    // pop stack[-1] until meet stack[-1]<current: previous stack[-1]>current, must > current stack[-1]
    stack:=[]int{}
    ans:=0
    for i:=0;i<=len(heights);i++{ // if col cannot meet its right boundary, use 0
        cur:=0
        if i<len(heights){
            cur=heights[i]
        }
        for len(stack)>0 && heights[stack[len(stack)-1]]>cur{
            h:=heights[stack[len(stack)-1]]
            stack=stack[:len(stack)-1]
            left:=-1 // if stack=[], means there is no lower col in left side
            if len(stack)>0{
                left=stack[len(stack)-1]
            }
            
            ans=max(ans,(i-left-1)*h)
        }
        stack=append(stack,i)
    }
    
    return ans
}