package array

func minOperations2654(nums []int) int {
	var l int = len(nums)
	var res int = 1<<31 - 1
	var one_cnt int = 0
	for i := 0; i < l; i++ {
		if nums[i] == 1 {
			one_cnt++
		}
	}
	if one_cnt > 0 {
		return l - one_cnt
	}
	for i := 0; i < l; i++ {
		common := nums[i]
		for j := i + 1; j < l; j++ {
			common = gcd(nums[j], common)
			if common == 1 {
				cur_steps := j - i //产生一个1需要的步骤数
				//rest_steps := l - (j - i + 1)
				res = min_int(res, cur_steps+l-1)
				break
			}
		}
	}
	if res == 1<<31-1 {
		return -1
	}
	return res
}
