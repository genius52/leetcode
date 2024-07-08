package array

// (sub[0] + sub[1]) % 2 == (sub[1] + sub[2]) % 2 == ... == (sub[x - 2] + sub[x - 1]) % 2
func MaximumLength3201(nums []int) int {
	var l int = len(nums)
	var pre_odd_one_cnt int = 0   //奇数结尾，余数为1的子序列长度
	var pre_odd_zero_cnt int = 0  //奇数结尾，余数为0的子序列长度
	var pre_even_one_cnt int = 0  //偶数结尾，余数为1的子序列长度
	var pre_even_zero_cnt int = 0 //偶数结尾，余数为0的子序列长度
	if nums[0]%2 == 0 {
		pre_even_one_cnt = 1
		pre_even_zero_cnt = 1
	} else {
		pre_odd_one_cnt = 1
		pre_odd_zero_cnt = 1
	}
	for i := 1; i < l; i++ {
		if nums[i]%2 == 0 { //偶数
			pre_even_zero_cnt++
			pre_even_one_cnt = max(pre_even_one_cnt, pre_odd_one_cnt+1)
		} else { //奇数
			pre_odd_zero_cnt++
			pre_odd_one_cnt = max(pre_odd_one_cnt, pre_even_one_cnt+1)
		}
	}
	return max(pre_odd_one_cnt, pre_odd_zero_cnt, pre_even_one_cnt, pre_even_zero_cnt)
}
