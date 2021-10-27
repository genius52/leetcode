
#1498. Number of Subsequences That Satisfy the Given Sum Condition
class Solution_1498:
    def numSubseq(self, nums: list[int], target: int) -> int:
        nums.sort()
        l = len(nums)
        left = 0
        right = l - 1
        res = 0
        while left <= right:
            if nums[left] + nums[right] <= target:
                res += 1<<(right-left)
                left += 1
            else:
                right -= 1
        return res % (10 ** 9 + 7)