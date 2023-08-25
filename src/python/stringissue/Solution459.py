class Solution459:
    def repeatedSubstringPattern(self, s: str) -> bool:
        s2 = s + s
        return s2[1:len(s2) - 1].find(s) != -1