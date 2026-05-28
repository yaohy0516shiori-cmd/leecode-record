class Solution:
    def myAtoi(self, s: str) -> int:
        i=0
        n=len(s)
        intmax=2**31-1
        inmin=-2**31
        while i<n and s[i]==' ':
            i+=1
        sign=1
        if i<n and s[i] in ["+","-"]:
            if s[i]=="-":
                sign=-1
            i+=1
        res=0
        while i<n and s[i].isdigit():
            digit=int(s[i])
            res=res*10+digit
            i+=1
        res*=sign
        if res<inmin:
            return inmin
        if res>intmax:
            return intmax
        return res
