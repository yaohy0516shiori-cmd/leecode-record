/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, targetSum int) int {
    res:=map[int]int{0:1}
    ans:=0
    var dfs func(node *TreeNode, prefix int)
    dfs=func(node *TreeNode, prefix int){
        if node==nil{
            return
        }
        prefix+=node.Val
        // res stores ancestor's sums, previous sums= one of ancestors' sum + target 
        ans+=res[prefix-targetSum]
        res[prefix]+=1
        dfs(node.Left,prefix)
        dfs(node.Right,prefix)
        res[prefix]-=1
    }
    dfs(root,0)
    return ans
}