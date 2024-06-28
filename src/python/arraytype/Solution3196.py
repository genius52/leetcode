#3196. Maximize Total Cost of Alternating Subarrays
from functools import cache
class Solution3196:
    def maximumTotalCost(self, nums: list[int]) -> int:
        @cache
        def dp_maximumTotalCost(idx: int,isflip: bool)->int:
            if idx == len(nums):
                return 0

            val = nums[idx]
            if isflip:
                val = -val
            ret1 = val + dp_maximumTotalCost(idx + 1,not isflip)
            ret2 = val + dp_maximumTotalCost(idx + 1,False)
            return max(ret1,ret2)

        return dp_maximumTotalCost(0, False)
