/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
    var prev *TreeNode
    var dfs func(node *TreeNode) bool 
    dfs=func(node *TreeNode) bool{
        // use inorder search
        if node==nil{ // nil is legal
            return true
        }
        if !dfs(node.Left){ // go to left, judge whether it is legal
            return false
        }
        if prev !=nil && node.Val<=prev.Val{ // when is not nil, judge the value
            return false
        }
        prev=node
        if !dfs(node.Right){ 
            return false
        }
        return true
    }
    return dfs(root)
}