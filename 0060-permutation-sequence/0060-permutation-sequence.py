class Solution:
    def getPermutation(self, n: int, k: int) -> str:
        nums=[str(i) for i in range(1,n+1)]
        fact=1
        for i in range(1,n):
            fact*=i
        '''
        For each position, if there are remaining numbers, fixing one number creates
        (remaining - 1)! permutations.
        One important boundary detail is that k is one-indexed in the problem statement,
        but array indices and group indices are zero-indexed.
        So I first convert k to zero-based indexing.
        Then for each position, I compute:

        index = k // group_size

        This tells me which remaining number should be placed at the current position.
        After choosing it, I remove that number from the list and update:

        k = k % group_size

        That means I continue searching inside the selected group.

        I repeat this until all positions are filled.

        Time complexity is O(n^2) because removing from a list costs O(n).
        Space complexity is O(n).
        '''
        k-=1
        ans=[]
        for size in range(n,0,-1):
            index=k//fact
            ans.append(nums.pop(index))
            k%=fact
            if size>1:
                fact//=(size-1)
        return "".join(ans)