import heapq
class MedianFinder:

    def __init__(self):
        self.small=[] # smallest heap, for n//2 max element
        self.large=[] # biggest heap, for n- n//2 min element

    def addNum(self, num: int) -> None:
        heapq.heappush(self.small,-num) # only have small heap in python
        if self.large and -self.small[0]>self.large[0]:
            val=-heapq.heappop(self.small)
            heapq.heappush(self.large,val)
        if len(self.small)>len(self.large)+1: # balance of two heaps, euqal to k=n//2
            val=-heapq.heappop(self.small)
            heapq.heappush(self.large,val)
        elif len(self.small)<len(self.large): 
            val=heapq.heappop(self.large)
            heapq.heappush(self.small,-val)


    def findMedian(self) -> float:
        if len(self.small)>len(self.large):
            return round(float(-self.small[0]),1)
        return round((-self.small[0]+self.large[0])/2.0,1)



# Your MedianFinder object will be instantiated and called as such:
# obj = MedianFinder()
# obj.addNum(num)
# param_2 = obj.findMedian()