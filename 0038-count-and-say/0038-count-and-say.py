class Solution:
    def countAndSay(self, n: int) -> str:
        if n==1:
            return '1'
        s='1'
        j=1
        while j<n:
            count=1
            temp=''
            for i in range(1,len(s)):
                if s[i] == s[i-1]:
                    count+=1
                else:
                    temp+=str(count)+s[i-1]
                    count=1
            temp+=str(count)+s[-1]
            s=temp
            j+=1
        return s