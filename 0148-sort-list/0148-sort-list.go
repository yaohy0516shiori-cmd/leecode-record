import "container/heap"
type item struct{
    node *ListNode
}
type minheap []item
func (h minheap) Len() int {return len(h)}
func (h minheap) Less(i,j int) bool{
    return h[i].node.Val<h[j].node.Val
}
func (h minheap) Swap(i,j int){
    h[i], h[j] = h[j], h[i]
}
func (h *minheap) Push(x any){
    *h=append(*h,x.(item))
}
func (h *minheap) Pop() any {
    old:=*h
    n:=len(old)
    x:=old[n-1]
    *h=old[:n-1]
    return x
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
    h:=&minheap{}
    heap.Init(h)
    cur:=head
    for cur!=nil{
        heap.Push(h,item{node: cur})
        cur=cur.Next
    }
    dummy := &ListNode{}
    cur = dummy
    for h.Len() > 0 {
        temp := heap.Pop(h).(item)
        cur.Next = temp.node
        cur = cur.Next
    }
    cur.Next = nil
    return dummy.Next
}