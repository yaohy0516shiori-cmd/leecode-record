class Solution:
    def reverse(self, x: int) -> int:
        res=0
        sign=1
        if x<0:
            sign=-1
            x=-x
        while x>0:
            temp=x%10
            x=x//10
            res=res*10+temp
        if res>2**31-1:
            return 0
        return res*sign