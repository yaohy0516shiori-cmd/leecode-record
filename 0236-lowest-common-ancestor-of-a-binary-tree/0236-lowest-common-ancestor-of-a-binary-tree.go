/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    // here, root also means node, once the node == q or p, return this node
    if root==nil{
        return nil
    }
    if root==p || root==q{
        return root
    }
    left:=lowestCommonAncestor(root.Left, p, q)
    right:=lowestCommonAncestor(root.Right, p, q)
    if left!=nil && right!=nil{
        return root // once p,q can be found in left/right tree--lowest common ancestor
    }
    if left!=nil{
        return left
    }
    if right!=nil{
        return right
    }
    return nil
}