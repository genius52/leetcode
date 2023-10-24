package array

func MinSizeSubarray(nums []int, target int) int {
	var l int = len(nums)
	var sum int = 0
	for i := 0; i < l; i++ {
		sum += nums[i]
	}
	if target%sum == 0 {
		return (target / sum) * l
	}
	var cnt int = target / sum * l
	target %= sum
	var prefix_sum map[int]int = make(map[int]int)
	var res int = 1<<31 - 1
	sum = 0
	prefix_sum[0] = -1
	for i := 0; i < l*2; i++ {
		sum += nums[i%l]
		need := sum - target
		if idx, ok := prefix_sum[need]; ok {
			gap := abs_int(i - idx)
			if gap+cnt < res {
				res = gap + cnt
			}
		}
		prefix_sum[sum] = i
	}
	if res == 1<<31-1 {
		return -1
	}
	return res
}
