class Solution:
    def permute(self, nums: List[int]) -> List[List[int]]:
        if len(nums)==1:
            return [nums]
        res=[]
        path=[]
        used=[False]*len(nums)
        def dfs():
            if len(path)==len(nums):
                res.append(path[:])
                return 
            for i in range(len(nums)):
                if used[i]:
                    continue
                path.append(nums[i])
                used[i]=True
                dfs()
                used[i]=False
                path.pop()
        dfs()
        return res