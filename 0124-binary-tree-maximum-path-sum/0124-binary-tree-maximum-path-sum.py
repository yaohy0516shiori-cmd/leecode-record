# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    ans=float('-inf') # as turning point
    def maxPathSum(self, root: Optional[TreeNode]) -> int:
        def dfs(node):
            if node is None:
                return 0 # None cannot be calculate
            left=max(dfs(node.left),0) # find all left node, start from the bottom
            right=max(dfs(node.right),0) # if contribution of this node value smaller than 0, we don't need it
            self.ans=max(left+right+node.val,self.ans)  # maximum result in the tree (left+right+node.val cannot return to top layer, illegal)
            return max(left+node.val,right+node.val,node.val) # return to top layer and continue
        dfs(root)
        return self.ans