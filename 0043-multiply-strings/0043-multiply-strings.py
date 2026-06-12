class Solution:
    def multiply(self, num1: str, num2: str) -> str:
        if num1=='0'or num2 =='0':
            return '0'
        n=len(num1)
        m=len(num2)
        res=[0]*(m+n)
        for i in range(n-1,-1,-1):
            for j in range(m-1,-1,-1):
                mul=int(num1[i])*int(num2[j])
                p1=i+j
                p2=i+j+1
                tol=mul+res[p2]
                res[p1]+=tol//10
                res[p2]=tol%10
        product=''.join(map(str,res)).lstrip('0')
        return product
        

