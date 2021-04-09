package array

func findMaxConsecutiveOnes(nums []int) int {
	nums = append(nums,0)
	var l int = len(nums)
	var left int = 0
	var max_len int = 0
	for left < l{
		if nums[left] == 0{
			left++
			continue
		}
		var right int = left + 1
		for right < l && nums[right] == 1{
			right++
		}
		cur_len := right - left
		if cur_len > max_len{
			max_len = cur_len
		}
		left = right
	}
	return max_len
}