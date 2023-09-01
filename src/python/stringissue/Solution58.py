class Solution58:
    def lengthOfLastWord(self, s: str) -> int:
        l = len(s)
        cnt = 0
        for i in range(l - 1,-1,-1):
            if s[i] != ' ':
                cnt += 1
            elif s[i] == ' ' and cnt > 0:
                return cnt
        return cnt