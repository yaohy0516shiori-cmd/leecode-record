class Solution:
    def threeSumClosest(self, nums: List[int], target: int) -> int:
        nums.sort()
        ans=nums[0]+nums[1]+nums[2]
        if len(nums)==3:
            return ans

        for i in range(len(nums)-2):
            left=i+1
            right=len(nums)-1
            while left<right:
                tol=nums[i]+nums[left]+nums[right]
                if abs(tol-target)<abs(ans-target):
                    ans=tol
                if tol==target:
                    return tol
                elif tol<target:
                    left+=1
                else:
                    right-=1
        return ans