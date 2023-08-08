class Solution389:
    def findTheDifference(self, s: str, t: str) -> str:
        record = [0 for _ in range(26)]
        for c in s:
            record[ord(c) - ord('a')] += 1
        for c in t:
            record[ord(c) - ord('a')] -= 1
            if record[ord(c) - ord('a')] < 0:
                return c
        return ''