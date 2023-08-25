class Solution283:
    def moveZeroes(self, nums: list[int]) -> None:
        """
        Do not return anything, modify nums in-place instead.
        """
        left = 0 # zero idx
        right = 0 # find no-zero number
        l = len(nums)
        while left < l and right < l:
            while left < l and nums[left] != 0:
                left += 1
            if left == l:
                break
            right = left + 1
            while right < l and nums[right] == 0:
                right += 1
            if right == l:
                break
            nums[left] = nums[right]
            nums[right] = 0

    def moveZeroes2(self, nums: list[int]) -> None:
        """
        Do not return anything, modify nums in-place instead.
        """
        l = len(nums)
        idx = 0
        for n in nums:
            if n != 0:
                nums[idx] = n
                idx += 1
        while idx < l:
            nums[idx] = 0
            idx += 1

