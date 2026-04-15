import "container/heap"
type Item struct{
    num int
    freq int
}
type minheap []Item
func (h minheap) Len() int {return len(h)}
// define this is small heap
func (h minheap) Less(i,j int) bool {
    return h[i].freq<h[j].freq
}
// provide swap func, same as big heap
func (h minheap) Swap(i,j int) {
    h[i],h[j]=h[j],h[i]
}
// append element to the last place, heap use it first then use 'up' function to resort the heap (truly maintain the heap structure )
func (h *minheap) Push(x any){
    *h=append(*h,x.(Item))
}
/*
provide a function to libarary function pop: heap.Pop() is different from h.Pop(), 
heap.Pop will exchange the first and last element
then resort 0-(n-2) element, then use h.Pop(), now we pop the first element of h
n := h.Len() - 1
h.Swap(0, n)
down(h, 0, n) move the new root to its right place
return h.Pop()
*/ 
func (h *minheap) Pop() any {
    old:=*h
    n:=len(old)
    x:=old[n-1]
    *h=old[:n-1]
    return x
}
func topKFrequent(nums []int, k int) []int {
    hashtable:=make(map[int]int)
    for _,i:=range nums{
        hashtable[i]++ 
    }
    h:= &minheap{}
    heap.Init(h)
    for num,freq:=range hashtable{
        heap.Push(h,Item{num:num,freq:freq})
        if h.Len()>k{
            heap.Pop(h)
        }
    }
    ans:=make([]int,k)
    for i:=0;i<len(ans);i++{
        ans[i]=heap.Pop(h).(Item).num
    }
    return ans
}