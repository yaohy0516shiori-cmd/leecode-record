/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
    new:=make(map[*Node]*Node)
    cur:=head
    for cur!=nil{
        new[cur]=&Node{Val:cur.Val}
        cur=cur.Next
    }
    cur=head
    for cur!=nil{
        new[cur].Next=new[cur.Next]
        new[cur].Random=new[cur.Random]
        cur=cur.Next
    }
    return new[head]
}