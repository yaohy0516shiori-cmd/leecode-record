# class Solution:
#     def trap(self, height: List[int]) -> int:
#         stack=[]
#         ans=0
#         n=len(height)
#         for i in range(n):
#             while stack and height[i]>height[stack[-1]]:
#                 top=stack.pop()
#                 if not stack:
#                     break
#                 left=stack[-1]
#                 w=i-left-1
#                 h=min(height[left],height[i])-height[top]
#                 ans+=w*h
#             stack.append(i)
#         return ans   

# from typing import List

# class Solution:
#     def trap(self, height: List[int]) -> int:
#         n = len(height)
# 每个位置能装水的量取决于左右两边最高的板的高度
#         pre_max = [0] * n
#         pre_max[0] = height[0]

#         for i in range(1, n):
#             pre_max[i] = max(pre_max[i - 1], height[i])

#         suf_max = [0] * n
#         suf_max[-1] = height[-1]

#         for i in range(n - 2, -1, -1):
#             suf_max[i] = max(suf_max[i + 1], height[i])

#         ans = 0

#         for h, pre, suf in zip(height, pre_max, suf_max):
#             ans += min(pre, suf) - h

#         return ans
# two pointer
class Solution:
    def trap(self,height: List[int]) -> int:
        # 简化前后缀数组方法本质, 还是找左右最大然后看当前位置的值
        n=len(height)
        ans=0
        l,r=0,n-1
        pre,suf=0,0
        while l<=r:
            pre=max(height[l],pre)
            suf=max(height[r],suf)
            if pre<suf:
                ans+=pre-height[l]
                l+=1
            else:
                ans+=suf-height[r]
                r-=1
        return ans