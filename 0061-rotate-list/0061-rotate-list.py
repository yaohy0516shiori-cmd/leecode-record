# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
'''
I first handle edge cases: if the list is empty, has only one node, or k is zero,
I return head directly.

Then I traverse the list once to compute its length and find the tail.
Since rotating by the list length gives the same list, I reduce k using modulo.

If the reduced k is zero, no rotation is needed.

Otherwise, I connect the tail to the head to form a circular linked list.
For a right rotation, the new tail is length minus k minus one steps from the
original head. The node after the new tail becomes the new head.

Finally, I break the circle by setting new_tail.next to None and return new_head.

The time complexity is linear in the number of nodes, because I traverse the list
to get the length and then move to the new tail.
The space complexity is constant, because I only use a few pointers.
'''
class Solution:
    def rotateRight(self, head: Optional[ListNode], k: int) -> Optional[ListNode]:
        # first, edge case is empty head or one Node or k equals 0
        if not head or not head.next or k==0:
            return head
        # second, try to make it into a circle listnode, so we can easily break it down when we find the new head
        tail=head
        length=1
        # third, counting node num
        while tail.next:
            tail=tail.next
            length+=1
        k%=length
        # reduce k using k modulo length because it rotated in the same listnode
        if k==0:
            return head
        # make it into a circle
        tail.next=head
        # find a new head. new tail position should be length minus k minus 1(no dummy node)
        step=length-k-1
        new=head
        for _ in range(step):
            new=new.next
        # then we move to that node and next node is new head
        newhead=new.next
        # break the circle
        new.next=None
        return newhead
