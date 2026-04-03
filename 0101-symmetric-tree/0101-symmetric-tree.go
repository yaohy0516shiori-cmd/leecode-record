/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
    if root==nil || (root!=nil && root.Left==nil && root.Right==nil){
        return true
    }
    var dfs func(a,b *TreeNode) bool
    dfs=func(a,b *TreeNode) bool{
        if a==nil && b==nil{
            return true
        }else if (a==nil && b!=nil) || (b==nil && a!=nil){
            return false
        }
        if a.Val!=b.Val{
            return false
        }
        hash:=dfs(a.Left,b.Right) && dfs(a.Right,b.Left)
        return hash
    }
    return dfs(root.Left,root.Right)
}