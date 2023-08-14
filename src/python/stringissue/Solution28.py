
#28. Find the Index of the First Occurrence in a String

class Solution28:
    def strStr(self, haystack: str, needle: str) -> int:
        l1 = len(haystack)
        l2 = len(needle)
        for i in range(l1):
            idx1 = i
            idx2 = 0
            while idx1 < l1 and idx2 < l2:
                if haystack[idx1] != needle[idx2]:
                    break
                idx1 += 1
                idx2 += 1
            if idx2 == l2:
                return i
        return -1