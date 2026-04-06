/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 // level order
// func rightSideView(root *TreeNode) []int {
//     if root==nil{
//         return []int{}
//     }
//     queue:=[]*TreeNode{root}
//     ans:=[]int{}
//     for len(queue)>0{
//         level:=[]int{}
//         size:=len(queue)
//         for j:=0;j<size;j++{
//             cur:=queue[0]
//             queue=queue[1:]
//             level=append(level,cur.Val)
//             if cur.Left!=nil{
//                 queue=append(queue,cur.Left)
//             }
//             if cur.Right!=nil{
//                 queue=append(queue,cur.Right)
//             }
//         }
//         ans=append(ans,level[len(level)-1])
//     }
//     return ans
// }

func rightSideView(root *TreeNode) []int {
    ans:=[]int{}
    var dfs func(node *TreeNode, depth int)
    dfs=func(node *TreeNode, depth int){
        if node==nil{
            return
        }
        // if ans length==depth, means it is the first element of this level
        if depth==len(ans){
            ans=append(ans,node.Val)
        }
        dfs(node.Right,depth+1)
        dfs(node.Left,depth+1)
        
    }
    dfs(root,0)
    return ans
}