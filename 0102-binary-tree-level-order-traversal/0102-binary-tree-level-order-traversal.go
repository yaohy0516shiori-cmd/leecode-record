/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 // level order
func levelOrder(root *TreeNode) [][]int {
    res:=[][]int{}
    if root==nil{
        return [][]int{}
    }
    queue:=[]*TreeNode{root}
    for len(queue)>0{
            level:=[]int{}
            size:=len(queue)
            for j:=0;j<size;j++{

                // pop the first element in queue
                cur:=queue[0]
                queue=queue[1:]
                level=append(level,cur.Val)

                // core code, use queue to store node from left to right
                if cur.Left!=nil{
                    queue=append(queue,cur.Left)
                }
                if cur.Right!=nil{
                    queue=append(queue,cur.Right)
                }
            }
            res=append(res,level) 
    }
    return res
}