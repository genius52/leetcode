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
        roman_int = {"I":1,"V":5,"X":10,"L":50,"C":100,"D":500,"M":1000}
        neighbour = {"V":"I","L":"X","D":"C","M":"D"}
        total = 0
        for i in range(l):
            if s[i] in roman_int:
                pass
        return total