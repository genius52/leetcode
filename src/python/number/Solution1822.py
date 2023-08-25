class Solution1822:
    def arraySign(self, nums: list[int]) -> int:
        positive = True
        for n in nums:
            if n == 0:
                return 0
            elif n < 0:
                positive = not positive

        if positive:
            return 1
        return -1