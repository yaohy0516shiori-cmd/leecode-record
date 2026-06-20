class Solution:
    def permuteUnique(self, nums: List[int]) -> List[List[int]]:
        nums.sort()
        res=[]
        if len(nums)==1:
            return [nums]
        used=[False]*len(nums)
        path=[]
        def dfs():
            if len(path)==len(nums):
                return res.append(path[:])
            for i in range(len(nums)):
                if used[i]:
                    continue
                if i>0 and nums[i]==nums[i-1] and used[i-1]:
                    continue
                used[i]=True
                path.append(nums[i])
                dfs()
                used[i]=False
                path.pop()
        dfs()
        return res