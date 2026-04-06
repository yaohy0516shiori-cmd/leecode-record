/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func flatten(root *TreeNode)  {
    // if root==nil || (root.Left==nil && root.Right==nil){
    //     return 
    // }
    // queue:=[]*TreeNode{}
    // var dfs func(node *TreeNode)
    // dfs=func(node *TreeNode){
    //     if node==nil{
    //         return
    //     }
    //     queue=append(queue,node)
    //     dfs(node.Left)
    //     dfs(node.Right)
    //     return
    // }
    // dfs(root)
    // for i:=0;i<len(queue);i++{
    //     queue[i].Left=nil
    //     if i!=len(queue)-1{
    //         queue[i].Right=queue[i+1]
    //     }
        
    // }
    var prev *TreeNode = nil
    var dfs func(node *TreeNode)
    dfs=func(node *TreeNode){
        if node==nil{
            return
        }
        dfs(node.Right)
        dfs(node.Left)
        node.Right=prev
        node.Left=nil
        prev=node
    }
    dfs(root)
}