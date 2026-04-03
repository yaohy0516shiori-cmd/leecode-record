/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum(root *TreeNode) int {
    ans:=root.Val
    var dfs func(node *TreeNode) int
    dfs=func(node *TreeNode) int {
        // bottom None is nil, return 0 to calculate, 1. avoid plus minus, 2. don't change ans
        if node == nil{
            return 0
        }
        l:=max(dfs(node.Left),0)
        r:=max(dfs(node.Right),0)
        ans=max(node.Val+l+r,ans)
        return max(node.Val+l,node.Val+r, node.Val)
    }
    dfs(root)
    return ans
}