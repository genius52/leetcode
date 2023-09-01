class Solution709:
    def toLowerCase(self, s: str) -> str:
        res = ""
        for c in s:
            if 'A' <= c <= 'Z':
                res += chr(ord(c) - ord('A') + ord('a'))
            else:
                res += c
        return res