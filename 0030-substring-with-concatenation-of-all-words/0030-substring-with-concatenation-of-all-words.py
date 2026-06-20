from collections import Counter
class Solution:
    def findSubstring(self, s: str, words: List[str]) -> List[int]:
        if not s or not words:
            return []
        word_len=len(words[0])
        total=word_len*len(words)
        w_count=len(words)
        target=Counter(words)
        res=[]

        for i in range(word_len):
            left=i
            curr=Counter()
            count=0
            for j in range(i,len(s)-word_len+1,word_len):
                word=s[j:j+word_len]
                if word in target:
                    curr[word]+=1
                    count+=1
                    while curr[word]>target[word]:
                        left_word=s[left:left+word_len]
                        curr[left_word]-=1
                        count-=1
                        left+=word_len
                    if count==w_count:
                        res.append(left)
                else:
                    curr.clear()
                    count=0
                    left=j+word_len
        return res