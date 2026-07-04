#Didnt solve
class Solution:
    def removeDuplicateLetters(self, s: str) -> str:
        cnt = Counter(s) 
        stack = deque()
        for c in s:
            if c not in stack:
                while len(stack) > 0 and ord(c) < ord(stack[-1]) and cnt[stack[-1]] > 0:
                    stack.pop()
                stack.append(c)
            cnt[c] -= 1
        return "".join(stack)
                
