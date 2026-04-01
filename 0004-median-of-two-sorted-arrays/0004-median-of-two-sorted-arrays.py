class Solution:
    def findMedianSortedArrays(self, nums1: List[int], nums2: List[int]) -> float:
        sums=len(nums1)+len(nums2)
        i=0
        j=0
        pre,cur=0,0 # doesn't need to store the whole array
        mid_place=int(sums/2)
        for k in range(mid_place+1):
            pre=cur # store the last one
            if i<len(nums1) and (j>=len(nums2) or nums1[i]<nums2[j]):
                cur=nums1[i]
                i+=1
            else:
                cur=nums2[j]
                j+=1
        if sums%2==0:
            return (pre+cur)/2
        else:
            return cur

