/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
    // tow vector must meet at first circle (before slow vector enter sencond circle)
    // therefore, 2*(a+b)=a+n(b+c)+b
    // a=c+(n-1)(b+c)
    fast,slow:=head,head
    for fast!=nil{
        slow=slow.Next
        if fast.Next==nil{
            return nil
        }
        fast=fast.Next.Next
        if fast==slow{
            p:=head
            for p!=slow{
                p=p.Next
                slow=slow.Next
            }
            return p
        }
    }
    return nil
}