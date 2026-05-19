# class Solution:
#     def decodeString(self, s: str) -> str:
#         def dfs(i):
#             res="" 这个res是每一层里的res要区分清楚或者直接在外层加个ans进来也行
#             while i<len(s):
#                 if s[i].isalpha():
#                     res+=s[i]
#                     i+=1
#                 elif s[i].isdigit():
#                     k=0
#                     while i<len(s) and s[i].isdigit():
#                         k=k*10+int(s[i])
#                         i+=1
#                     i+=1
#                     inner,i=dfs(i) 返回res和复制次数
#                     res+=inner*k
#                 elif s[i]==']':
#                     return res,i+1
#             return res,i
#         ans,_=dfs(0)
#         return ans

class Solution:
    def decodeString(self, s: str) -> str:
        stack = []
        curr_str = ""
        curr_num = 0

        for ch in s:
            if ch.isdigit():
                curr_num = curr_num * 10 + int(ch)
            # stack 存外层状态+循环次数
            elif ch == "[":
                stack.append((curr_str, curr_num))
                curr_str = ""
                curr_num = 0
            # 这里读到了开始改curstr
            elif ch == "]":
                prev_str, num = stack.pop()
                curr_str = prev_str + curr_str * num

            else:
                curr_str += ch

        return curr_str