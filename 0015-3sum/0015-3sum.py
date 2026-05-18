class Solution:
    def threeSum(self, nums: list[int]) -> list[list[int]]:
        n=len(nums)
        if n==3:
            return [nums] if nums[0]+nums[1]+nums[2]==0 else []
        res=[]
        nums.sort()
        for i in range(n-2):
            if nums[i]>0:
                break
            if i>0 and nums[i-1]==nums[i]:
                continue
            l,r=i+1,n-1
            while l<r:
                ans=nums[l]+nums[i]+nums[r]
                if ans<0:
                    l+=1
                elif ans>0:
                    r-=1
                else:
                    res.append([nums[l],nums[i],nums[r]])
                    l+=1
                    r-=1
                    while l<r and nums[l-1]==nums[l]:
                        l+=1
                    while l<r and nums[r+1]==nums[r]:
                        r-=1
        return res