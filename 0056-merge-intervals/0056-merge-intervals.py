class Solution:
    def merge(self, intervals: List[List[int]]) -> List[List[int]]:
        intervals.sort(key=lambda x:x[0])
        n=len(intervals)
        res=[]
        start=intervals[0][0]
        end=intervals[0][1]
        for i in range(1,n):
            if end<intervals[i][0]:
                res.append([start,end])
                start=intervals[i][0]
                end=intervals[i][1]
                continue
            else:
                end=max(end,intervals[i][1])
        res.append([start,end])
        return res

# class Solution:
#     def merge(self, intervals: List[List[int]]) -> List[List[int]]:
#         intervals.sort(key=lambda x: x[0])

#         res = []

#         for start, end in intervals:
#             if not res or res[-1][1] < start:
#                 res.append([start, end])
#             else:
#                 res[-1][1] = max(res[-1][1], end)

#         return res