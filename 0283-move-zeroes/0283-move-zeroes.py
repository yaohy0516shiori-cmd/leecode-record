class Solution:
    def moveZeroes(self, nums: List[int]) -> None:
        """
        Do not return anything, modify nums in-place instead.
        """
        n=len(nums)
        i = 0

        for j in range(len(nums)):
            if nums[j] != 0:
                nums[i], nums[j] = nums[j], nums[i]
                i += 1
    # def moveZeroes(self, nums: List[int]) -> None:
    #     n = len(nums)

    #     # i 先找到第一个 0
    #     i = 0
    #     while i < n and nums[i] != 0:
    #         i += 1

    #     # j 从第一个 0 后面开始找非 0
    #     j = i + 1

    #     while j < n:
    #         if nums[j] != 0:
    #             nums[i], nums[j] = nums[j], nums[i]
    #             i += 1
    #         j += 1
