import "math"
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    if len(nums1)>len(nums2){
        nums1,nums2=nums2,nums1
    }
    total:=len(nums1)+len(nums2)
    k:=(total+1)/2
    m,n:=len(nums1),len(nums2)
    left,right:=0,m
    for left<=right{
        i:=(left+right)/2
        j:=k-i
        aleft:=math.Inf(-1)
		if i > 0 {
			aleft = float64(nums1[i-1])
		}
        aright := math.Inf(1)
		if i < m {
			aright = float64(nums1[i])
		}

		bleft := math.Inf(-1)
		if j > 0 {
			bleft = float64(nums2[j-1])
		}

		bright := math.Inf(1)
		if j < n {
			bright = float64(nums2[j])
		}

        if aleft<=bright && bleft<aright{
            if (m+n)%2==0{
                return (max(aleft,bleft)+min(aright,bright))/2
            } else{ 
                return max(aleft,bleft)
            } 
        }
        if aleft>bright{ 
            right=i-1
        } else{ //use left window 
            left=i+1
        } //use right window
    }
    return 0
}