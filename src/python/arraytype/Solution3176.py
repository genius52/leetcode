#3176. Find the Maximum Length of a Good Subsequence I
from functools import cache
class Solution3176:
    def maximumLength(self, nums: list[int], k: int) -> int:
        @cache
        def dp_maximumLength(idx: int, k: int, cnt: int, last: int) -> int:
            if cnt > k:
                return -(1 << 31)
            if idx == len(nums):
                return 0
            if nums[idx] == last:
                return 1 + dp_maximumLength(idx + 1, k, cnt, last)
            else:
                if last == 0:
                    ret1 = 1 + dp_maximumLength(idx + 1, k, cnt, nums[idx])
                else:
                    ret1 = 1 + dp_maximumLength(idx + 1, k, cnt + 1, nums[idx])
                ret2 = dp_maximumLength(idx + 1, k, cnt, last)  # skip
                return max(ret1, ret2)

        res = dp_maximumLength(0,k,0,0)
        dp_maximumLength.cache_clear()
        return res