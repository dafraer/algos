class Solution:
    def calculate(self, s: str) -> int:
        tokens = []
        mult = []
        i = 0
        for j in range(len(s)):
            if s[j] in ("+", "-", "*", "/"):
                tokens.append(int(s[i:j].strip()))
                tokens.append(s[j])
                i = j + 1
            elif j == len(s) - 1:
                tokens.append(int(s[i : j + 1].strip()))
        if len(tokens) == 1:
            return tokens[0]

        res = tokens[0]
        for i in range(1, len(tokens) - 1, 2):
            if tokens[i] == "*":
                res *= tokens[i + 1]
            elif tokens[i] == "/":
                res //= tokens[i + 1]
            else:
                mult.append(res)
                mult.append(tokens[i])
                res = tokens[i + 1]
        mult.append(res)
