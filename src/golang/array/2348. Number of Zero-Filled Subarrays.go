package array

func zeroFilledSubarray(nums []int) int64 {
	var l int = len(nums)
	var res int64 = 0
	var left int = 0
	for left < l{
		for left < l && nums[left] != 0{
			left++
		}
		if left == l{
			break
		}
		var right int = left
		for right < l && nums[right] == 0{
			res += int64(right - left + 1)
			right++
		}
		left = right + 1
	}
	return res
}