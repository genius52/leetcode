# Symbol       Value
# I             1
# V             5
# X             10
# L             50
# C             100
# D             500
# M             1000

class Solution13:
    def romanToInt(self, s: str) -> int:
        l = len(s)
        val = {"I":1,"V":5,"X":10,"L":50,"C":100,"D":500,"M":1000}
        sequence = {"I": 0,"V":1,"X":2,"L":3,"C":4,"D":5,"D":6,"M":7}
        sum = 0
        for i in range(l):
            if i + 1 < l and sequence[s[i]] < sequence[s[i + 1]]:
                sum -= val[s[i]]
            else:
                sum += val[s[i]]
        return sum