class Solution:
    def isMatch(self, s: str, p: str) -> bool:
        i,j=0,0  # i is s pointer, j is p
        start=-1
        mat=-1
        while i<len(s):
            # normal situation, move forward
            if j<len(p) and (p[j]==s[i] or p[j]=="?"):
                i+=1
                j+=1
            # meet * then j move forward
            elif j<len(p) and p[j]=='*':
                start=j
                mat=i
                j+=1
            # if not match, go back to the place then j plus 1
            elif start !=-1:
                j=start+1
                mat+=1
                i=mat
            else:
                return False
        # if s is over, rest of p must equal *
        while j<len(p) and p[j]=='*':
            j+=1
        return j==len(p)
