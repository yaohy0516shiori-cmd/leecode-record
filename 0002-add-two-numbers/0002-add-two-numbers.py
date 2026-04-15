# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def addTwoNumbers(self, l1: Optional[ListNode], l2: Optional[ListNode]) -> Optional[ListNode]:
        dummy=ListNode()
        cur=dummy
        mark=0
        while l1 or l2 or mark:
            x = l1.val if l1 else 0
            y = l2.val if l2 else 0

            temp = x + y + mark
            cur.next = ListNode(temp % 10)
            mark = temp // 10

            cur = cur.next

            if l1:
                l1 = l1.next
            if l2:
                l2 = l2.next

        return dummy.next