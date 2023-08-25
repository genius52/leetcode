
class Solution:
    def plusOne(self, digits: list[int]) -> list[int]:
        l = len(digits)
        upgrade = True
        for i in range(l - 1, -1, -1):
            if upgrade:
                digits[i] += 1
                if digits[i] == 10:
                    digits[i] = 0
                    upgrade = True
                else:
                    upgrade = False
            else:
                break
        if upgrade:
            digits = [1] + digits
        return digits