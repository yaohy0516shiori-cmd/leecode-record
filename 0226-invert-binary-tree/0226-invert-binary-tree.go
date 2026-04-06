/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
    var dfs func(node *TreeNode)
    dfs=func(node *TreeNode){
        if node==nil{
            return
        }
        temp:=node.Left
        node.Left=node.Right
        node.Right=temp
        dfs(node.Left)
        dfs(node.Right)
    }
    dfs(root)
    return root
}