package array

func LongestSubarray2419(nums []int) int {
	var l int = len(nums)
	var max_val int = 1
	var max_len int = 0
	for i := 0;i < l;{
		if nums[i] < max_val{
			i++
			continue
		}
		if nums[i] > max_val{
			max_len = 1
		}
		max_val = nums[i]
		j := i
		for ;j < l && nums[i] == nums[j];j++{
			cur_len := j - i + 1
			if cur_len > max_len{
				max_len = cur_len
			}
		}
		i = j
	}
	return max_len
}