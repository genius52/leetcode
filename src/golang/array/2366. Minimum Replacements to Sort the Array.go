package array

func MinimumReplacement(nums []int) int64 {
	var l int = len(nums)
	var res int64 = 0
	var pre int = nums[l - 1]
	for i := l - 2;i >= 0;i--{
		if nums[i] <= pre{
			pre = nums[i]
		}else{
			cnt := (nums[i] + pre - 1)/ pre //10,3
			res += int64(cnt - 1)
			pre = nums[i] / cnt
		}
	}
	return res
}