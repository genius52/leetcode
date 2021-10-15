package array

func KLengthApart(nums []int, k int) bool {
	var l int = len(nums)
	var left int = 0
	for left < l{
		for left < l && nums[left] != 1{
			left++
		}
		if left >= l{
			break
		}
		var right int = left + 1
		for right < l && nums[right] != 1{
			right++
		}
		if right >= l{
			break
		}
		distance := right - left
		if distance <= k{
			return false
		}
		left = right
	}
	return true
}