class Solution:
    def divide(self, dividend: int, divisor: int) -> int:
        in_max=2**31-1
        in_min=-2**31
        # 特殊溢出情况
        if dividend == in_min and divisor == -1:
            return in_max

        # 判断符号
        sign = -1 if (dividend < 0) ^ (divisor < 0) else 1

        rest=abs(dividend)
        num=abs(divisor)
        ans=0
        while rest>=num:
            temp=num
            count=1
            while rest>=(temp<<1):
                temp<<=1
                count<<=1
            rest-=temp
            ans+=count

        if sign==-1:
            ans=-ans
        return ans
