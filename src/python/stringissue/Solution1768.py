class Solution1768:
    def mergeAlternately(self, word1: str, word2: str) -> str:
        l1 = len(word1)
        l2 = len(word2)
        idx1 = 0
        idx2 = 0
        res = ""
        while idx1 < l1 or idx2 < l2:
            if idx1 < l1:
                res += word1[idx1]
                idx1 += 1
            if idx2 < l2:
                res += word2[idx2]
                idx2 += 1
        return res
