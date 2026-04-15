# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def removeNthFromEnd(self, head: Optional[ListNode], n: int) -> Optional[ListNode]:
        dummy=ListNode(0,head)
        cur=dummy
        count=0
        node=head
        while node:
            count+=1
            node=node.next
        
        for i in range(count-n):
            cur=cur.next
        cur.next=cur.next.next
        return dummy.next