/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
    // inorder list change into a tree, -1 <= left tree node - right tree node <=1
    l,r:=0,len(nums)-1

    var build func(left,right int)*TreeNode
    build=func(left,right int)*TreeNode{
        if left>right{
            return nil
        }
        mid:=(left+right)/2
        rootval:=nums[mid]
        root:=&TreeNode{Val:rootval}
        root.Left=build(left,mid-1)
        root.Right=build(mid+1,right)

        return root
    }
    return build(l,r)
}