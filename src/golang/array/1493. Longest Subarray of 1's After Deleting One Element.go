package array

func LongestSubarray(nums []int) int {
	var l int = len(nums)
	var left int = 0
	var right int = 0
	var zero_cnt int = 0
	var res int = 0
	for left < l && right < l{
		for right < l{
			if nums[right] == 0{
				zero_cnt++
				if zero_cnt > 1{
					right++
					break
				}
			}
			res = max_int(res,right - left  + 1 - zero_cnt)
			right++
		}

		for left < l{
			if nums[left] == 0{
				left++
				zero_cnt--
				break
			}
			left++
		}
	}
	if res == l{
		return res - 1
	}
	return res
}