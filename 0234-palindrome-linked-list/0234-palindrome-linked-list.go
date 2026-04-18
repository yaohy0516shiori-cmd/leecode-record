/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
    if head== nil || head.Next==nil{
        return true
    }
    slow,fast:=head,head
    for fast!=nil && fast.Next!=nil{
        fast=fast.Next.Next
        slow=slow.Next
    }
    if fast!=nil{
        slow=slow.Next
    }
    second:=reverse(slow)
    p1,p2:=head,second
    for p2!=nil{
        if p1.Val!=p2.Val{
            return false
        }
        p1=p1.Next
        p2=p2.Next
    }
    return true
}

func reverse(head *ListNode) *ListNode {
    var prev *ListNode
    cur:=head
    for cur !=nil{
        next:=cur.Next
        cur.Next=prev
        prev=cur
        cur=next
    }
    return prev
}