type MinStack struct {
    minus []int
    stack []int
}


func Constructor() MinStack {
    return MinStack{
        minus: []int{},
        stack: []int{},
    }
}


func (this *MinStack) Push(val int)  {
    this.stack=append(this.stack,val)
    if len(this.minus)!=0{
        this.minus=append(this.minus,min(this.minus[len(this.minus)-1],val))
    }else{
        this.minus=append(this.minus,val)
    }
}


func (this *MinStack) Pop()  {
    this.stack=this.stack[:len(this.stack)-1]
    this.minus=this.minus[:len(this.minus)-1]
}


func (this *MinStack) Top() int {
    return this.stack[len(this.stack)-1]
}


func (this *MinStack) GetMin() int {
    return this.minus[len(this.minus)-1]
}



/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */