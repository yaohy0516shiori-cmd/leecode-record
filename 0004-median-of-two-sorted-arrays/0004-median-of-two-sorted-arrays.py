class Solution:
    def findMedianSortedArrays(self, a: List[int], b: List[int]) -> float:
        # O(m+n)
        '''sums=len(nums1)+len(nums2)
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
            return cur'''
        if len(a)>len(b):
            a,b=b,a
        m,n=len(a),len(b)
        left,right=0,m
        half=(m+n+1)//2 # odd: (m+n+1)//2, not index, even: (m+n)//2==(m+n+1)//2
        while left<=right:
            i=(left+right)//2
            j=half-i

            aleft=a[i-1] if i>0 else float('-inf')
            aright=a[i]  if i<m else float('inf')
            bleft=b[j-1] if j>0 else float('-inf')
            bright=b[j]  if j<n else float('inf')

            if aleft<=bright and bleft<=aright:
                if not (m+n)%2:
                    return (max(aleft,bleft)+min(aright,bright))/2
                else: 
                    return max(aleft,bleft)
            if aleft>bright:
                right=i-1 # use left window
            else:
                left=i+1 # use right window
