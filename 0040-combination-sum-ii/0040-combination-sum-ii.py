class Solution:
    def combinationSum2(self, candidates: List[int], target: int) -> List[List[int]]:
        res=[]
        ans=[]
        candidates.sort()
        n=len(candidates)
        def dfs(index,target):
            if target==0:
                res.append(ans[:])
                return
            for i in range(index,n):
                if i>index and candidates[i]==candidates[i-1]:
                    continue
                if target<candidates[i]:
                    break
                ans.append(candidates[i])
                dfs(i+1,target-candidates[i])
                ans.pop()
        dfs(0,target)
        return res
