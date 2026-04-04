/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
    index:=make(map[int]int)
    // use hash to record the position, inorder element's left is its left tree
    for i:=0;i<len(inorder);i++{
        index[inorder[i]]=i
    }
    var build func(inleft,inright,preleft,preright int) *TreeNode
    build=func(inleft,inright,preleft,preright int) *TreeNode{
        if inleft>inright || preleft>preright{
            return nil
        }
        // root is the first in preorder
        root_val:=preorder[preleft]
        root:=&TreeNode{Val:root_val}
        // root in inorder position
        id:=index[root_val]
        // left tree node number
        left:=id-inleft
        // inleft do not need change coz is the left side of left tree, id-1 is right side of left tree in inorder
        // preleft+1 always is the root of child, rightside is preleft+left in preorder
        root.Left=build(inleft,id-1,preleft+1,preleft+left)
        // id+1 is left side in in order, inright do not need to be changed
        // preleft+left+1 is left side of right tree in preorder, preright do not need change
        root.Right=build(id+1,inright,preleft+1+left,preright)

        return root
    }
    return build(0,len(inorder)-1,0,len(preorder)-1)

}