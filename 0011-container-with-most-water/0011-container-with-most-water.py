class Solution:
    def maxArea(self, height: List[int]) -> int:
        left,right=0,len(height)-1
        if len(height)==2:
            return min(height[0],height[1])
        
        ans=min(height[0],height[1])
        while left<right:
            current=(right-left)*min(height[left],height[right])
            ans=max(ans,current)
            if height[left]<height[right]:
                left+=1
            else:
                right-=1
        
        return ans


