# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def swapPairs(self, head: Optional[ListNode]) -> Optional[ListNode]:
        dummy=ListNode()
        dummy.next=head
        cur=dummy
        while cur.next and cur.next.next:
            temp=cur.next
            second=temp.next
            temp.next=second.next
            second.next=temp
            cur.next=second
            cur=cur.next.next
        return dummy.next