class Solution1502:
    def canMakeArithmeticProgression(self, arr: list[int]) -> bool:
        l = len(arr)
        if l <= 1:
            return True
        arr = sorted(arr)
        for i in range(1,l - 1):
            if (arr[i] - arr[i - 1]) != (arr[i + 1] - arr[i]):
                return False
        return True