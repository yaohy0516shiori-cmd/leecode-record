class Solution:
    def nextPermutation(self, nums: List[int]) -> None:
        """
        Do not return anything, modify nums in-place instead.
        从右往左找第一个 nums[i] < nums[i+1] 的位置，说明 nums[i] 可以被右边某个更大的数替换，从而让整体变大一点。
        然后从右往左找第一个 nums[j] > nums[i] 的数，交换 nums[i] 和 nums[j]。
        交换后，为了让结果尽量小，需要把 i 后面的后缀反转成升序。
        如果一开始找不到 nums[i] < nums[i+1]，说明整个数组是递减的，已经是最大排列，直接反转成最小排列。
        """
        n=len(nums)
        i=n-2
        # 倒序找第一个升序位置
        while i>=0 and nums[i]>=nums[i+1]:
            i-=1
        if i>=0:
            j=n-1
            # 找第一个比nums[i]大的数
            while nums[j]<=nums[i]:
                j-=1
            nums[i],nums[j]=nums[j],nums[i]
        left=i+1
        right=n-1
        # 倒转列表
        while left<right:
            nums[left],nums[right]=nums[right],nums[left]
            left+=1
            right-=1