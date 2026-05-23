class Solution:
    def productExceptSelf(self, nums: List[int]) -> List[int]:
        ans=[1 for _ in range(len(nums))]
        cur=1
        # pre product,(product before i)
        for i in range(len(nums)):
            ans[i]*=cur
            cur*=nums[i]
        res=1
        # suf product
        for i in range(len(nums)-1,-1,-1):
            ans[i]*=res
            res*=nums[i]
        return ans