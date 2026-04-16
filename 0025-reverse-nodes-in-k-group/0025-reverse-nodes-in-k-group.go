/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
    dummy:=&ListNode{Val:0,Next:head}
    prev:=dummy
    for{
        // kth element of this group
        kth:=getend(prev,k)
        if kth==nil{
            break
        }
        // next group start position
        nextgroup:=kth.Next
        // reverse the group[cur:start]
        start:=nextgroup
        cur:=prev.Next
        for cur!=nextgroup{
            temp:=cur.Next
            cur.Next=start
            start=cur
            cur=temp
        }
        // prev.next=old cur, now in new end
        old:=prev.Next
        // old prev connect to kth(new start)
        prev.Next=kth
        // prev jump to new start
        prev=old
    }
    return dummy.Next
}

func getend(cur *ListNode, k int) *ListNode{
    for cur!=nil && k>0{
        cur=cur.Next
        k--
    }
    return cur
}