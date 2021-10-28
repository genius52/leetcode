
#1498. Number of Subsequences That Satisfy the Given Sum Condition
# For each A[i], find out the maximum A[j]
# that A[i] + A[j] <= target.
#
# For each elements in the subarray A[i+1] ~ A[j],
# we can pick or not pick,
# so there are 2 ^ (j - i) subsequences in total.
# So we can update res = (res + 2 ^ (j - i)) % mod.

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