class Solution896:
    def isMonotonic(self, nums: list[int]) -> bool:
        l = len(nums)
        find_diff = False
        increase = False
        for i in range(1,l):
            diff1 = nums[i] - nums[i - 1]
            if find_diff:
                if (increase and diff1 < 0) or (not increase and diff1 > 0):
                    return False
            else:
                if diff1 > 0:
                    increase = True
                    find_diff = True
                elif diff1 < 0:
                    increase = False
                    find_diff = True
        return True