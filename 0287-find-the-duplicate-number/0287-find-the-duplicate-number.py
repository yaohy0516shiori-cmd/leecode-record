from typing import List

class Solution:
    def findDuplicate(self, nums: List[int]) -> int:
        slow = nums[0]
        fast = nums[0]

        # 第一阶段：快慢指针在环内相遇
        while True:
            # 重复值成环，模拟链表
            slow = nums[slow]
            fast = nums[nums[fast]]
            # 相遇
            if slow == fast:
                break

        # 第二阶段：找环入口
        slow = nums[0]

        while slow != fast:
            slow = nums[slow]
            fast = nums[fast]

        return slow