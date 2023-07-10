package array

func longestAlternatingSubarray(nums []int, threshold int) int {
	var l int = len(nums)
	var pre int = 0
	var res int = 0
	for i := l - 1; i >= 0; i-- {
		if nums[i] > threshold {
			pre = 0
			continue
		}
		if i == (l - 1) {
			pre = 1
			if (nums[i]%2) == 0 && pre > res {
				res = pre
			}
			continue
		}
		if (nums[i] % 2) == 0 {
			if (nums[i+1] % 2) != 0 {
				pre++
			} else {
				pre = 1
			}
			if pre > res {
				res = pre
			}
		} else {
			if (nums[i+1] % 2) == 0 {
				pre++
			} else {
				pre = 1
			}
		}
	}
	return res
}
