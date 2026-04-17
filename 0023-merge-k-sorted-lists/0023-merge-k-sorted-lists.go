import "container/heap"
type Item struct{
    node *ListNode
}
type minheap []Item
func (h minheap) Len() int {return len(h)}
// define this is small heap
func (h minheap) Less(i,j int) bool {
    return h[i].node.Val<h[j].node.Val
}
// provide swap func, same as big heap
func (h minheap) Swap(i,j int) {
    h[i],h[j]=h[j],h[i]
}
// append element to the last place, heap use it first then use 'up' function to resort the heap (truly maintain the heap structure )
func (h *minheap) Push(x any){
    *h=append(*h,x.(Item))
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
func mergeKLists(lists []*ListNode) *ListNode {
    dummy:=&ListNode{Val:0}
    h:=&minheap{}
    heap.Init(h)
    for _,i := range lists{
        if i!=nil{
            heap.Push(h,Item{node:i})
        }
    }
    cur:=dummy
    for h.Len()>0{
        item:=heap.Pop(h).(Item)
        node:=item.node
        cur.Next=node
        cur=cur.Next
        if node.Next!=nil{
            heap.Push(h,Item{node:node.Next})
        }
    }
    return dummy.Next
}