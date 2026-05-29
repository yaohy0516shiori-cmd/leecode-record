class Solution:
    def isPalindrome(self, x: int) -> bool:
        if x<0:
            return False
        res=0
        org=x
        while org:
            temp=org%10
            res=res*10+temp
            org=org//10
        return res==x