package array

func MaximumScore(nums []int, k int) int {
	var l int = len(nums)
	var res int = 0
	var left int = k
	var right int = k
	var min_val int = nums[k]
	for left >= 0 || right < l {
		for left >= 0 && nums[left] >= min_val {
			left--
		}
		for right < l && nums[right] >= min_val {
			right++
		}
		res = max_int(res, min_val*(right-left-1))
		if left < 0 && right >= l {
			break
		}
		if left < 0 {
			min_val = nums[right]
		} else if right >= l {
			min_val = nums[left]
		} else {
			if nums[left] < nums[right] {
				min_val = nums[right]
			} else {
				min_val = nums[left]
			}
		}
	}
	return res
}
