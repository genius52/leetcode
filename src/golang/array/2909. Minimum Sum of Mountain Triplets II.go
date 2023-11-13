package array

func minimumSum2(nums []int) int {
	var l int = len(nums)
	var prefix []int = make([]int, l)
	var suffix []int = make([]int, l)
	prefix[0] = nums[0]
	for i := 1; i < l; i++ {
		prefix[i] = min_int(prefix[i-1], nums[i])
	}
	suffix[l-1] = nums[l-1]
	for i := l - 2; i >= 0; i-- {
		suffix[i] = min_int(suffix[i+1], nums[i])
	}
	var res int = 1<<31 - 1
	for i := 1; i < l-1; i++ {
		if prefix[i-1] < nums[i] && suffix[i+1] < nums[i] {
			sum := prefix[i-1] + nums[i] + suffix[i+1]
			if sum < res {
				res = sum
			}
		}
	}
	if res == 1<<31-1 {
		return -1
	}
	return res
}
