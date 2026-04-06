/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
    count:=0
    ans:=0
    var dfs func(node *TreeNode)
    dfs=func(node *TreeNode){
        if node==nil{
            return
        }
        dfs(node.Left)
        count++
        if count==k{
            ans=node.Val
            return
        }
        dfs(node.Right)
    }
    dfs(root)
    return ans
}