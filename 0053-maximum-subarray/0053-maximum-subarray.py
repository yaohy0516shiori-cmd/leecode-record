from typing import List

class Solution:
    def maxSubArray(self, nums: List[int]) -> int:
        prefix_sum = 0
        min_prefix = 0
        ans = nums[0]

        for num in nums:
            prefix_sum += num
            # 一定是减去前面出现过的最小前缀和，不是全局最大减所有最小,会出现不存在的数组
            ans = max(ans, prefix_sum - min_prefix)

            min_prefix = min(min_prefix, prefix_sum)

        return ans

# DP
# from typing import List

# class Solution:
#     def maxSubArray(self, nums: List[int]) -> int:
#         cur = nums[0]
#         ans = nums[0]

#         for i in range(1, len(nums)):
#             cur = max(nums[i], cur + nums[i])
#             ans = max(ans, cur)

#         return ans