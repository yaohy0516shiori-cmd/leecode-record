class Solution:
    def subarraySum(self, nums: List[int], k: int) -> int:
        res=defaultdict(int)
        prefix=0
        ans=0
        for i in range(len(nums)):
            # 一定要有res[0],否则prefix=k无法计入
            res[prefix]+=1
            prefix+=nums[i]
            ans+=res[prefix-k]
        return ans
        