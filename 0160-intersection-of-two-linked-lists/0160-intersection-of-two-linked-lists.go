/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    hash:=make(map[*ListNode]int)
    cur:=headA
    for cur != nil{
        hash[cur]=1
        cur=cur.Next
    }
    cur=headB
    for cur != nil{
        if hash[cur]==1{
            return cur
        }
        cur=cur.Next
    }
    return nil
}