class Solution:
    def longestConsecutive(self, nums: List[int]) -> int:
        hashmap=defaultdict()
        longest=0
        for i in nums:
            if i in hashmap:
                hashmap[i]+=1
            else:
                hashmap[i]=1
        for key in hashmap:
            if key-1 not in hashmap:
                count=1
                x=key
                while x+1 in hashmap:
                    count+=1
                    x+=1
                longest=max(count,longest)
        return longest
