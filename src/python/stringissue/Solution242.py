class Solution242:
    def isAnagram(self, s: str, t: str) -> bool:
        letter = [0] * 26
        for i in range(len(s)):
            letter[ord(s[i]) - ord('a')] += 1
        for i in range(len(t)):
            letter[ord(t[i]) - ord('a')] -= 1
        for i in range(26):
            if letter[i] != 0:
                return False
        return True
