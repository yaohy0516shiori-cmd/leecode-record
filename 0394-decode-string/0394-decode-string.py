class Solution:
    def decodeString(self, s: str) -> str:
        def dfs(i):
            res=""
            while i<len(s):
                if s[i].isalpha():
                    res+=s[i]
                    i+=1
                elif s[i].isdigit():
                    k=0
                    while i<len(s) and s[i].isdigit():
                        k=k*10+int(s[i])
                        i+=1
                    i+=1
                    inner,i=dfs(i)
                    res+=inner*k
                elif s[i]==']':
                    return res,i+1
            return res,i
        ans,_=dfs(0)
        return ans