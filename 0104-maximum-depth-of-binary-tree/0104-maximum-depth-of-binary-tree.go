/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
    depth:=0
    if root==nil{
        return depth
    }
    queue:=[]*TreeNode{root}
    for len(queue)>0{
        depth+=1
        size:=len(queue)
        for j:=0;j<size;j++{
            cur:=queue[0]
            queue=queue[1:]
            if cur.Left!=nil{
                queue=append(queue,cur.Left)
            }
            if cur.Right != nil{
                queue=append(queue,cur.Right)
            }
        }
    }
    return depth
}