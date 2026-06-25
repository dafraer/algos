class Solution:
    def getHint(self, secret: str, guess: str) -> str:
        b, c = 0, 0
        cnt = Counter(guess)

        e = [True for _ in secret]
        for i in range(len(secret)):
            if secret[i] == guess[i]:
                cnt[secret[i]] -= 1
                e[i] = False
                b += 1
        
        for i in range(len(secret)):
            if cnt[secret[i]] > 0 and e[i]:
                cnt[secret[i]] -= 1
                c += 1

        return f'{b}A{c}B'